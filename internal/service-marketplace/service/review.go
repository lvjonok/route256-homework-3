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

func (s *Service) AddReview(ctx context.Context, req *pb.AddReviewRequest) (*pb.AddReviewResponse, error) {
	newReview := models.Review{
		ProductID: types.ID(req.ProductID),
		Text:      req.Text,
	}

	id, err := s.DB.CreateReview(ctx, &newReview)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to add review, err: <%v>", err)
	}

	return &pb.AddReviewResponse{ID: uint64(*id)}, nil
}

func (s *Service) GetReviews(ctx context.Context, req *pb.GetReviewsRequest) (*pb.GetReviewsResponse, error) {
	res, err := s.DB.GetProductReviews(ctx, types.Int2ID(req.ProductID))
	if err != nil {
		if err == repo.ErrNotFound {
			return nil, status.Errorf(codes.NotFound, "there are no reviews")
		}

		return nil, status.Errorf(codes.Internal, "failed to get reviews for product, err: <%v>", err)
	}

	pbReviews := []*pb.Review{}
	for _, r := range res {
		pbReviews = append(pbReviews, &pb.Review{
			ID:        uint64(r.ID),
			ProductID: uint64(r.ProductID),
			Text:      r.Text,
		})
	}

	return &pb.GetReviewsResponse{Reviews: pbReviews}, nil
}
