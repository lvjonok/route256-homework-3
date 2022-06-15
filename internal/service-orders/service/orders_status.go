package service

import (
	"context"

	types "gitlab.ozon.dev/lvjonok/homework-3/core/models"
	"gitlab.ozon.dev/lvjonok/homework-3/internal/service-orders/models"
	api "gitlab.ozon.dev/lvjonok/homework-3/pkg/srv_orders/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

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
