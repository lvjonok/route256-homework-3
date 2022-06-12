package service

import (
	"context"

	"gitlab.ozon.dev/lvjonok/homework-3/core/cacheconnector"
	types "gitlab.ozon.dev/lvjonok/homework-3/core/models"
	"gitlab.ozon.dev/lvjonok/homework-3/internal/service-marketplace/models"
	"gitlab.ozon.dev/lvjonok/homework-3/internal/service-marketplace/repo"
	pb "gitlab.ozon.dev/lvjonok/homework-3/pkg/srv_marketplace/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	s.Metrics.CreateProductInc()
	newProduct := models.Product{
		Name: req.Name,
		Desc: req.Desc,
	}

	id, err := s.DB.CreateProduct(ctx, &newProduct)
	if err != nil {
		s.Metrics.CreateProductErrorsInc()

		return nil, status.Errorf(codes.Internal, "failed to create new product, err: <%v>", err)
	}

	return &pb.CreateProductResponse{ID: uint64(*id)}, nil
}

func (s *Service) GetProduct(ctx context.Context, req *pb.GetProductRequest) (*pb.GetProductResponse, error) {
	s.Metrics.GetProductInc()

	var (
		product *models.Product
		err     error
	)

	if product, err = s.Cache.GetProduct(ctx, *types.Int2ID(req.GetID())); err == cacheconnector.ErrCacheMiss {
		product, err = s.DB.GetProduct(ctx, types.Int2ID(req.ID))
		if err != nil {
			if err == repo.ErrNotFound {
				return nil, status.Errorf(codes.NotFound, "there is no product")
			}

			s.Metrics.GetProductErrorsInc()
			return nil, status.Errorf(codes.Internal, "failed to get product, err: <%v>", err)
		}

		if err := s.Cache.UpsertProduct(ctx, *product); err != nil {
			return nil, status.Errorf(codes.Internal, "failed to upsert product in cache, err: <%v>", err)
		}
	} else if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to check product in cache, err: <%v>", err)
	}

	return &pb.GetProductResponse{
		ID:   uint64(product.ID),
		Name: product.Name,
		Desc: product.Desc,
	}, nil
}
