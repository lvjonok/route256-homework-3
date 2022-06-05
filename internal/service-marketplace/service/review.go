package service

import (
	"context"

	pb "gitlab.ozon.dev/lvjonok/homework-3/pkg/srv_marketplace/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) AddReview(ctx context.Context, req *pb.AddReviewRequest) (*pb.AddReviewResponse, error) {

	return nil, status.Error(codes.Unimplemented, "haha")
}

func (s *Service) GetReviews(ctx context.Context, req *pb.GetReviewsRequest) (*pb.GetReviewsResponse, error) {

	return nil, status.Error(codes.Unimplemented, "haha")
}
