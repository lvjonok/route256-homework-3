package service

import (
	"context"

	types "gitlab.ozon.dev/lvjonok/homework-3/core/models"
	"gitlab.ozon.dev/lvjonok/homework-3/internal/service-warehouse/models"
	common "gitlab.ozon.dev/lvjonok/homework-3/pkg/common/api"
	marketplace_api "gitlab.ozon.dev/lvjonok/homework-3/pkg/srv_marketplace/api"
	"gitlab.ozon.dev/lvjonok/homework-3/pkg/srv_warehouse/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// RegisterProduct checks that product is registered on the marketplace and adds new entry to warehouse
func (s *Service) RegisterProduct(ctx context.Context, req *api.RegisterProductRequest) (*api.RegisterProductResponse, error) {
	s.Metrics.RegisterProductInc()

	// check that product is registered in marketplace
	s.Log.Debug("calling get product from marketplace client")
	if _, err := s.mpClient.GetProduct(ctx, &marketplace_api.GetProductRequest{
		ID: req.ProductID,
	}); err != nil {
		if st, ok := status.FromError(err); ok && st.Code() == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "there is no such product in marketplace, err: <%v>", err)
		}

		s.Metrics.RegisterProductErrorsInc()
		return nil, status.Errorf(codes.Internal, "failed to check that product exists, err: <%v>", err)
	}

	// adding new entry
	if err := s.DB.RegisterProduct(ctx, &models.Entry{
		ProductID: types.ID(req.ProductID),
		Quantity:  int(req.Quantity),
	}); err != nil {
		s.Metrics.RegisterProductErrorsInc()
		return nil, status.Errorf(codes.Internal, "failed to register product, err: <%v>", err)
	}

	return &api.RegisterProductResponse{ProductID: req.ProductID}, nil
}

// CheckProducts returns current quantity of each requests product in the warehouse
func (s *Service) CheckProducts(ctx context.Context, req *api.CheckProductsRequest) (*api.CheckProductsResponse, error) {
	s.Metrics.CheckProductsInc()

	ids := []types.ID{}
	for _, i := range req.ProductIDs {
		ids = append(ids, *types.Int2ID(i))
	}
	units, err := s.DB.CheckProducts(ctx, ids)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to query check products, err: <%v>", err)
	}

	pUnits := []*common.ProductUnit{}
	for _, u := range units {
		pUnits = append(pUnits, &common.ProductUnit{
			ProductID: uint64(u.ProductID),
			Quantity:  uint64(u.Quantity),
		})
	}

	return &api.CheckProductsResponse{
		Units: pUnits,
	}, nil
}

// BookProducts locks requested products from warehouse to be used in order and returns ids of entries with booking
func (s *Service) BookProducts(ctx context.Context, req *api.BookProductsRequest) (*api.BookProductsResponse, error) {
	s.Metrics.BookProductsInc()

	units := []types.ProductUnit{}
	for _, u := range req.Units {
		units = append(units, types.ProductUnit{
			ProductID: types.ID(u.ProductID),
			Quantity:  int(u.Quantity),
		})
	}

	bookingIDs, err := s.DB.BookProducts(ctx, units)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to book products, err: <%v>", err)
	}

	ids := []uint64{}
	for _, i := range bookingIDs {
		ids = append(ids, uint64(i))
	}

	return &api.BookProductsResponse{BookingIDs: ids}, nil
}

// UnbookProducts accepts entries we want to unlock and releases them
func (s *Service) UnbookProducts(ctx context.Context, req *api.UnbookProductsRequest) (*common.Empty, error) {
	s.Metrics.UnbookProductsInc()

	ids := []types.ID{}
	for _, id := range req.BookingIDs {
		ids = append(ids, types.ID(id))
	}

	if err := s.DB.UnbookProducts(ctx, ids); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to unbook products, err: <%v>")
	}

	return &common.Empty{}, nil
}
