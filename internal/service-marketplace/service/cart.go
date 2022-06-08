package service

import (
	"context"

	types "gitlab.ozon.dev/lvjonok/homework-3/core/models"
	"gitlab.ozon.dev/lvjonok/homework-3/internal/service-marketplace/models"
	"gitlab.ozon.dev/lvjonok/homework-3/internal/service-marketplace/repo"
	pb "gitlab.ozon.dev/lvjonok/homework-3/pkg/srv_marketplace/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) UpdateCart(ctx context.Context, req *pb.UpdateCartRequest) (*pb.UpdateCartResponse, error) {
	s.Metrics.UpdateCartInc()

	products := []models.ProductUnit{}
	for _, r := range req.Products {
		products = append(products, models.ProductUnit{
			ID:       types.ID(r.ProductID),
			Quantity: int(r.Quantity),
		})
	}

	newCart := models.Cart{
		UserID:   types.ID(req.ID),
		Products: products,
	}

	id, err := s.DB.UpdateCart(ctx, &newCart)
	if err != nil {
		if err == repo.ErrNotFound {
			return nil, status.Errorf(codes.NotFound, "there is no cart")
		}

		s.Metrics.UpdateCartErrorsInc()
		return nil, status.Errorf(codes.Internal, "failed to update cart, err: <%v>", err)
	}

	return &pb.UpdateCartResponse{ID: uint64(*id)}, nil
}

func (s *Service) GetCart(ctx context.Context, req *pb.GetCartRequest) (*pb.GetCartResponse, error) {
	s.Metrics.GetCartInc()

	cart, err := s.DB.GetCart(ctx, types.Int2ID(req.ID))
	if err != nil {
		if err == repo.ErrNotFound {
			return nil, status.Errorf(codes.NotFound, "there is no cart")
		}

		s.Metrics.GetCartErrorsInc()
		return nil, status.Errorf(codes.Internal, "failed to get cart, err: <%v>", err)
	}

	pbProducts := []*pb.ProductUnit{}
	for _, p := range cart.Products {
		pbProducts = append(pbProducts, &pb.ProductUnit{
			ProductID: uint64(p.ID),
			Quantity:  uint64(p.Quantity),
		})
	}

	return &pb.GetCartResponse{ID: uint64(cart.UserID), Products: pbProducts}, nil
}
