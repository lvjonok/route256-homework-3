package service

import (
	"context"
	"errors"

	"gitlab.ozon.dev/lvjonok/homework-3/core/cacheconnector"
	types "gitlab.ozon.dev/lvjonok/homework-3/core/models"
	"gitlab.ozon.dev/lvjonok/homework-3/internal/service-marketplace/models"
	"gitlab.ozon.dev/lvjonok/homework-3/internal/service-marketplace/repo"
	pb "gitlab.ozon.dev/lvjonok/homework-3/pkg/srv_marketplace/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) AddReview(ctx context.Context, req *pb.AddReviewRequest) (*pb.AddReviewResponse, error) {
	s.Metrics.AddReviewInc()

	newReview := models.Review{
		ProductID: types.ID(req.ProductID),
		Text:      req.Text,
	}

	id, err := s.DB.CreateReview(ctx, &newReview)
	if err != nil {
		s.Metrics.AddReviewErrorsInc()

		return nil, status.Errorf(codes.Internal, "failed to add review, err: <%v>", err)
	}

	if err := s.Cache.AppendReview(ctx, newReview); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to append review in cache, err: <%v>", err)
	}

	return &pb.AddReviewResponse{ID: uint64(*id)}, nil
}

func (s *Service) GetReviews(ctx context.Context, req *pb.GetReviewsRequest) (*pb.GetReviewsResponse, error) {
	s.Metrics.GetReviewsInc()

	var (
		res []models.Review
		err error
	)

	if res, err = s.Cache.GetReviews(ctx, *types.Int2ID(req.ProductID)); errors.Is(err, cacheconnector.ErrCacheMiss) {
		res, err = s.DB.GetProductReviews(ctx, types.Int2ID(req.ProductID))
		if err != nil {
			if err == repo.ErrNotFound {
				return nil, status.Errorf(codes.NotFound, "there are no reviews")
			}

			s.Metrics.GetReviewsErrorsInc()
			return nil, status.Errorf(codes.Internal, "failed to get reviews for product, err: <%v>", err)
		}

		if err := s.Cache.UpsertReviews(ctx, *types.Int2ID(req.ProductID), res); err != nil {
			return nil, status.Errorf(codes.Internal, "failed to upsert reviews in cache, err: <%v>", err)
		}
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
