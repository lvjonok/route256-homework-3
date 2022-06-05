package service

import (
	"context"

	pb "gitlab.ozon.dev/lvjonok/homework-3/pkg/srv_marketplace/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {

	return nil, status.Error(codes.Unimplemented, "haha")
}

func (s *Service) GetProduct(ctx context.Context, req *pb.GetProductRequest) (*pb.GetProductResponse, error) {

	return nil, status.Error(codes.Unimplemented, "haha")
}

func (s *Service) UpdateCart(ctx context.Context, req *pb.UpdateCartRequest) (*pb.UpdateCartResponse, error) {

	return nil, status.Error(codes.Unimplemented, "haha")
}
