package service_test

import (
	"context"
	"testing"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	types "gitlab.ozon.dev/lvjonok/homework-3/core/models"
	"gitlab.ozon.dev/lvjonok/homework-3/internal/service-marketplace/models"
	"gitlab.ozon.dev/lvjonok/homework-3/internal/service-marketplace/service"
	service_marketplace "gitlab.ozon.dev/lvjonok/homework-3/pkg/srv_marketplace/api"
	"go.uber.org/zap"
)

func TestAddReview(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	logger, err := zap.NewDevelopment()
	require.NoError(t, err)

	mockDB := service.NewDBMock(mc)
	mockDB.CreateReviewMock.Return(types.Int2ID(1234), nil)

	mockMetrics := service.NewMetricsMock(mc)
	mockMetrics.AddReviewIncMock.Return()

	srv := service.New(mockDB, mockMetrics, logger)
	ctx := context.Background()

	resp, err := srv.AddReview(ctx, &service_marketplace.AddReviewRequest{
		ProductID: 1234,
		Text:      "text",
	})
	require.NoError(t, err)
	require.Equal(t, uint64(1234), resp.ID)
}

func TestGetReviews(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	logger, err := zap.NewDevelopment()
	require.NoError(t, err)

	mockDB := service.NewDBMock(mc)
	mockDB.GetProductReviewsMock.Return([]models.Review{
		{ID: 1, ProductID: 1234, Text: "review1"},
		{ID: 2, ProductID: 1234, Text: "review2"},
	}, nil)

	mockMetrics := service.NewMetricsMock(mc)
	mockMetrics.GetReviewsIncMock.Return()

	srv := service.New(mockDB, mockMetrics, logger)
	ctx := context.Background()

	resp, err := srv.GetReviews(ctx, &service_marketplace.GetReviewsRequest{
		ProductID: 1234,
	})
	require.NoError(t, err)
	require.Len(t, resp.Reviews, 2)

	require.Equal(t, uint64(1), resp.Reviews[0].ID)
	require.Equal(t, uint64(1234), resp.Reviews[0].ProductID)
	require.Equal(t, "review1", resp.Reviews[0].Text)

	require.Equal(t, uint64(2), resp.Reviews[1].ID)
	require.Equal(t, uint64(1234), resp.Reviews[1].ProductID)
	require.Equal(t, "review2", resp.Reviews[1].Text)
}
