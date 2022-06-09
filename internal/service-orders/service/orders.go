package service

import (
	"context"
	"time"

	types "gitlab.ozon.dev/lvjonok/homework-3/core/models"
	"gitlab.ozon.dev/lvjonok/homework-3/internal/service-orders/models"
	common "gitlab.ozon.dev/lvjonok/homework-3/pkg/common/api"
	service_marketplace "gitlab.ozon.dev/lvjonok/homework-3/pkg/srv_marketplace/api"
	api "gitlab.ozon.dev/lvjonok/homework-3/pkg/srv_orders/api"
	service_warehouse "gitlab.ozon.dev/lvjonok/homework-3/pkg/srv_warehouse/api"
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

	// get current existing product units from warehouse
	productIds := []uint64{}
	for _, r := range cart.Products {
		productIds = append(productIds, r.ProductID)
	}
	checkResp, err := s.whClient.CheckProducts(ctx, &service_warehouse.CheckProductsRequest{
		ProductIDs: productIds,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get existing products from warehouse, err: <%v>", err)
	}
	existing := map[uint64]uint64{}
	for _, el := range checkResp.Units {
		existing[el.ProductID] = el.Quantity
	}

	// check that we have enough items
	for _, want := range cart.Products {
		if amount, ok := existing[want.ProductID]; !ok || amount < want.Quantity {
			return nil, status.Errorf(codes.FailedPrecondition, "there are not enough products in the warehouse want: <%v>, exists: <%v>", want.Quantity, amount)
		}
	}

	// book items in warehouse
	bookingResp, err := s.whClient.BookProducts(ctx, &service_warehouse.BookProductsRequest{Units: cart.Products})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to book products in the warehouse, err: <%v>", err)
	}

	units := []types.ProductUnit{}
	for _, r := range cart.Products {
		units = append(units, types.ProductUnit{
			ProductID: types.ID(r.ProductID),
			Quantity:  int(r.Quantity),
		})
	}
	id, err := s.DB.CreateOrder(ctx, &models.Order{
		UserID:   types.ID(req.UserID),
		Products: units,
		Status:   "created",
	})

	if err != nil {
		// as we did not succeeded in order creation, we want to unlock items in warehouse for other people
		go func(units []*common.ProductUnit) {
			if _, err := s.whClient.UnbookProducts(ctx, &service_warehouse.UnbookProductsRequest{BookingIDs: bookingResp.BookingIDs}); err != nil {
				// delay between calls
				time.Sleep(100 * time.Millisecond)
				s.Log.Error("failed to unbook products, retrying in 100ms", zap.Error(err))
			}
			// inlocking items in warehouse
		}(cart.Products)

		s.Metrics.CreateOrderErrorsInc()
		return nil, status.Errorf(codes.Internal, "failed to create a new order")
	}

	return &api.CreateOrderResponse{OrderID: uint64(*id)}, nil
}

// CheckStatus simply returns status of asked order
func (s *Service) CheckStatus(ctx context.Context, req *api.CheckStatusRequest) (*api.CheckStatusResponse, error) {
	s.Metrics.CheckStatusInc()

	orderStatus, err := s.DB.CheckStatus(ctx, types.Int2ID(req.OrderID))
	if err != nil {
		s.Metrics.CheckStatusErrorsInc()
		return nil, status.Errorf(codes.Internal, "failed to get status of order, err: <%v>", err)
	}

	return &api.CheckStatusResponse{Status: orderStatus}, nil
}

// UpdateStatus sets new status for asked order
func (s *Service) UpdateStatus(ctx context.Context, req *api.UpdateStatusRequest) (*api.UpdateStatusResponse, error) {
	s.Metrics.UpdateStatusInc()

	if err := s.DB.UpdateStatus(ctx, &models.Order{
		OrderID: types.ID(req.OrderID),
		Status:  req.Status,
	}); err != nil {
		s.Metrics.UpdateStatusErrorsInc()
		return nil, status.Errorf(codes.Internal, "failed to update status of order, err: <%v>", err)
	}

	return &api.UpdateStatusResponse{OrderID: req.OrderID}, nil
}
