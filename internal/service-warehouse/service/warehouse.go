package service

import (
	"context"

	"gitlab.ozon.dev/lvjonok/homework-3/core/cacheconnector"
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
	if _, err := s.mpClient.GetProduct(ctx, &marketplace_api.GetProductRequest{
		ID: req.ProductID,
	}); err != nil {
		if st, ok := status.FromError(err); ok && st.Code() == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "there is no such product in marketplace, err: <%v>", err)
		}

		s.Metrics.RegisterProductErrorsInc()
		return nil, status.Errorf(codes.Internal, "failed to check that product exists, err: <%v>", err)
	}

	// as we updated product, we should delete it from cache, to fetch next time
	if err := s.Cache.DeleteProducts(ctx, []types.ID{types.ID(req.ProductID)}); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete product from cache, err: <%v>", err)
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

	var (
		units []types.ProductUnit
		err   error
	)

	// check in cache
	units, err = s.Cache.GetProducts(ctx, ids)
	if err == cacheconnector.ErrCacheMiss || len(units) != len(ids) {
		// we need to find indices which were in request, but did not appear in units

		// create hash table of ids we already found
		reqIds := map[types.ID]bool{}
		for _, i := range units {
			reqIds[i.ProductID] = true
		}

		missingIds := []types.ID{}
		// iterate through requests ids to find missing
		for _, i := range ids {
			// check if we did not find
			if _, ok := reqIds[i]; !ok {
				missingIds = append(missingIds, i)
			}
		}

		dbunits, err := s.DB.CheckProducts(ctx, missingIds)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to query check products, err: <%v>", err)
		}

		// update cache
		if err := s.Cache.UpsertProducts(ctx, dbunits); err != nil {
			return nil, status.Errorf(codes.Internal, "failed to update product units in cache, err: <%v>", err)
		}

		units = append(units, dbunits...)
	} else if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get products from cache: <%v>", err)
	}

	foundIds := map[uint64]bool{}

	pUnits := []*common.ProductUnit{}
	for _, u := range units {
		foundIds[uint64(u.ProductID)] = true

		pUnits = append(pUnits, &common.ProductUnit{
			ProductID: uint64(u.ProductID),
			Quantity:  uint64(u.Quantity),
		})
	}

	for _, id := range req.ProductIDs {
		if _, ok := foundIds[id]; !ok {
			pUnits = append(pUnits, &common.ProductUnit{
				ProductID: id,
				Quantity:  0,
			})
		}
	}

	return &api.CheckProductsResponse{
		Units: pUnits,
	}, nil
}

// BookProducts locks requested products from warehouse to be used in order and returns ids of entries with booking
func (s *Service) BookProducts(ctx context.Context, req *api.BookProductsRequest) (*api.BookProductsResponse, error) {
	s.Metrics.BookProductsInc()

	updatedIds := []types.ID{}
	units := []types.ProductUnit{}
	for _, u := range req.Units {
		updatedIds = append(updatedIds, *types.Int2ID(u.ProductID))
		units = append(units, types.ProductUnit{
			ProductID: types.ID(u.ProductID),
			Quantity:  int(u.Quantity),
		})
	}

	if err := s.Cache.DeleteProducts(ctx, updatedIds); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete products from cache, err: <%v>", err)
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

	// we should remove products from cache, to fetch them later
	if err := s.Cache.DeleteProducts(ctx, ids); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete products from cache, err: <%v>", err)
	}

	if err := s.DB.UnbookProducts(ctx, ids); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to unbook products, err: <%v>", err)
	}

	return &common.Empty{}, nil
}
