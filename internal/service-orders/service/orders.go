package service

import (
	"context"
	"time"

	types "gitlab.ozon.dev/lvjonok/homework-3/core/models"
	"gitlab.ozon.dev/lvjonok/homework-3/internal/service-orders/config"
	"gitlab.ozon.dev/lvjonok/homework-3/internal/service-orders/models"
	"gitlab.ozon.dev/lvjonok/homework-3/internal/service-orders/repo"
	service_marketplace "gitlab.ozon.dev/lvjonok/homework-3/pkg/srv_marketplace/api"
	api "gitlab.ozon.dev/lvjonok/homework-3/pkg/srv_orders/api"
	"go.uber.org/zap"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CreateOrder gets current cart of user, checks existence of products in warehouse microservice, books required amount, and, finally, creates a new order
func (s *Service) CreateOrder(ctx context.Context, req *api.CreateOrderRequest) (*api.CreateOrderResponse, error) {
	s.Metrics.CreateOrderInc()

	// getting cart of user
	cart, err := s.mpClient.GetCart(ctx, &service_marketplace.GetCartRequest{ID: req.UserID})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get cart of user, err: <%v>", err)
	}

	units := []types.ProductUnit{}
	for _, r := range cart.Products {
		units = append(units, types.ProductUnit{
			ProductID: types.ID(r.ProductID),
			Quantity:  int(r.Quantity),
		})
	}
	id, err := s.DB.CreateOrder(ctx, &models.Order{
		UserID:     types.ID(req.UserID),
		Products:   units,
		Status:     "created",
		SagaStatus: models.Created,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create ")
	}

	endedStatus := models.Created

	// start saga
	saga, err := s.ProcessOrder(ctx, &api.ProcessOrderRequest{OrderID: uint64(*id)})
	if err != nil {
		s.Log.Sugar().Errorf("failed to proccess order, err: <%v>", err)
	}
	if saga != nil {
		endedStatus = models.OrderStatus(saga.LastStatus)
	}

	return &api.CreateOrderResponse{OrderID: uint64(*id), LastStatus: string(endedStatus)}, nil
}

// ProcessOrder is orchestrating saga, accepting OrderID we want to process, it gets it and tries to run from the last step
func (s *Service) ProcessOrder(ctx context.Context, req *api.ProcessOrderRequest) (*api.ProcessOrderResponse, error) {
	order, err := s.DB.GetOrder(ctx, types.Int2ID(req.OrderID))
	if err != nil && err == repo.ErrNotFound {
		return nil, status.Errorf(codes.NotFound, "there is no order we could process")
	} else if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get order to process, err: <%v>", err)
	}

	if order.SagaStatus == models.Created {
		if err := s.handleCreated(ctx, order); err != nil {
			// nothing to be compensated
			go func(ctx context.Context, orderID types.ID) {
				for i := 0; i < 5; i++ {
					if reterr := s.DB.AddRetry(ctx, orderID, models.Created); reterr != nil {
						s.Log.Sugar().Errorf("failed to add retry, err: <%v>", reterr)
						time.Sleep(10 * time.Second)
					} else {
						// successfully added
						return
					}
				}
			}(context.Background(), order.OrderID)

			return nil, err
		}
	}

	if order.SagaStatus == models.Checked {
		if err := s.handleChecked(ctx, order); err != nil {
			// we either booked and it was successfully, or there is nothing to be compensated
			go func(ctx context.Context, orderID types.ID) {
				for i := 0; i < 5; i++ {
					if reterr := s.DB.AddRetry(ctx, orderID, models.Checked); reterr != nil {
						s.Log.Sugar().Errorf("failed to add retry, err: <%v>", reterr)
						time.Sleep(10 * time.Second)
					} else {
						// successfully added
						return
					}
				}
			}(context.Background(), order.OrderID)
			return nil, err
		}
	}

	// we do not do anything after this
	if order.SagaStatus == models.Booked {
		return &api.ProcessOrderResponse{LastStatus: string(order.SagaStatus)}, nil
	}

	return &api.ProcessOrderResponse{LastStatus: string(order.SagaStatus)}, nil
}

// SagaWorker performs `saga cleanup` process, by trying to finish orders which are not in the final state
func (s *Service) SagaWorker(ctx context.Context, cfg config.SagaConfig) {
	for {
		s.Log.Debug("starting saga worker loop")

		ids, err := s.DB.GetProcessingOrders(ctx, cfg.Retries)
		if err != nil {
			s.Log.Sugar().Error(status.Errorf(codes.Internal, "failed to get processing orders, err: <%v>", err))
		}

		for _, i := range ids {
			s.Log.Debug("processing order", zap.Uint("id", uint(i)))

			sagastatus, serr := s.ProcessOrder(ctx, &api.ProcessOrderRequest{OrderID: uint64(i)})
			s.Log.Debug("processing result", zap.Any("saga status", sagastatus), zap.Error(serr))

			if serr != nil {
				s.Log.Sugar().Error(status.Errorf(codes.Internal, "failed to process order <%v>, err: <%v>", i, err))
			}
			if sagastatus != nil {
				if err := s.DB.AddRetry(ctx, i, models.OrderStatus(sagastatus.LastStatus)); err != nil {
					s.Log.Sugar().Error("failed to add retry for order, err: <%v>", err)
				}
			}
		}

		s.Log.Debug("saga worker sleep")
		time.Sleep(time.Duration(cfg.TimeoutMs) * time.Millisecond)
	}
}
