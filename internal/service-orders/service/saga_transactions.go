package service

import (
	"context"

	"gitlab.ozon.dev/lvjonok/homework-3/internal/service-orders/models"
	common "gitlab.ozon.dev/lvjonok/homework-3/pkg/common/api"
	service_warehouse "gitlab.ozon.dev/lvjonok/homework-3/pkg/srv_warehouse/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// On the stage `created` we want to chech that there are enough items in the warehouse
func (s *Service) handleCreated(ctx context.Context, order *models.Order) error {
	productIds := []uint64{}
	for _, r := range order.Products {
		productIds = append(productIds, uint64(r.ProductID))
	}
	checkResp, err := s.whClient.CheckProducts(ctx, &service_warehouse.CheckProductsRequest{
		ProductIDs: productIds,
	})
	if err != nil {
		return status.Errorf(codes.Internal, "failed to get existing products from warehouse, err: <%v>", err)
	}
	existing := map[uint64]uint64{}
	for _, el := range checkResp.Units {
		existing[el.ProductID] = el.Quantity
	}

	// check that we have enough items
	for _, want := range order.Products {
		if amount, ok := existing[uint64(want.ProductID)]; !ok || amount < uint64(want.Quantity) {
			return status.Errorf(codes.FailedPrecondition, "there are not enough products in the warehouse want: <%v>, exists: <%v>", want.Quantity, amount)
		}
	}

	order.SagaStatus = models.Checked
	if err := s.DB.UpdateOrderSagaStatus(ctx, order); err != nil {
		return status.Errorf(codes.Internal, "failed to update order saga status, err: <%v>", err)
	}

	return nil
}

// On the stage `checked` we want to book items from the warehouse binding to our order
func (s *Service) handleChecked(ctx context.Context, order *models.Order) error {
	pbUnits := []*common.ProductUnit{}
	for _, u := range order.Products {
		pbUnits = append(pbUnits, &common.ProductUnit{
			ProductID: uint64(u.ProductID),
			Quantity:  uint64(u.Quantity),
		})
	}

	// book items in warehouse
	_, err := s.whClient.BookProducts(ctx, &service_warehouse.BookProductsRequest{Units: pbUnits})
	if err != nil {
		return status.Errorf(codes.Internal, "failed to book products in the warehouse, err: <%v>", err)
	}

	order.SagaStatus = models.Booked
	if err := s.DB.UpdateOrderSagaStatus(ctx, order); err != nil {
		return status.Errorf(codes.Internal, "failed to update order saga status, err: <%v>", err)
	}

	return nil
}
