package service_test

import (
	"context"
	"testing"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	types "gitlab.ozon.dev/lvjonok/homework-3/core/models"
	"gitlab.ozon.dev/lvjonok/homework-3/internal/service-orders/models"
	"gitlab.ozon.dev/lvjonok/homework-3/internal/service-orders/service"
	common "gitlab.ozon.dev/lvjonok/homework-3/pkg/common/api"
	marketplace_api "gitlab.ozon.dev/lvjonok/homework-3/pkg/srv_marketplace/api"
	api "gitlab.ozon.dev/lvjonok/homework-3/pkg/srv_orders/api"
	warehouse_api "gitlab.ozon.dev/lvjonok/homework-3/pkg/srv_warehouse/api"
	"go.uber.org/zap"
)

func TestCreateOrder(t *testing.T) {
	ctx := context.Background()

	mc := minimock.NewController(t)
	defer mc.Finish()

	logger, err := zap.NewDevelopment()
	require.NoError(t, err)

	mockMetrics := service.NewMetricsMock(mc)
	mockMetrics.CreateOrderIncMock.Return()

	mockDB := service.NewDBMock(mc)
	mockDB.CreateOrderMock.Expect(ctx, &models.Order{
		UserID: 100,
		Products: []types.ProductUnit{
			{
				ProductID: 1,
				Quantity:  2,
			},
			{
				ProductID: 2,
				Quantity:  4,
			},
		},
		Status: "created",
	})
	mockDB.CreateOrderMock.Return(types.Int2ID(0), nil)

	mockMarketplaceClient := service.NewMarketplaceClientMock(mc)
	mockMarketplaceClient.GetCartMock.Return(&marketplace_api.GetCartResponse{
		ID: 100,
		Products: []*common.ProductUnit{
			{
				ProductID: 1,
				Quantity:  2,
			},
			{
				ProductID: 2,
				Quantity:  4,
			},
		},
	}, nil)

	mockWarehouseClient := service.NewWarehouseClientMock(mc)
	mockWarehouseClient.CheckProductsMock.Return(&warehouse_api.CheckProductsResponse{
		Units: []*common.ProductUnit{
			{
				ProductID: 1,
				Quantity:  2,
			},
			{
				ProductID: 2,
				Quantity:  4,
			},
		},
	}, nil)
	mockWarehouseClient.BookProductsMock.Return(&warehouse_api.BookProductsResponse{
		BookingIDs: []uint64{0, 1},
	}, nil)

	srv := service.New(mockDB, mockMetrics, logger, mockMarketplaceClient, mockWarehouseClient)

	resp, err := srv.CreateOrder(ctx, &api.CreateOrderRequest{
		UserID: 100,
	})
	require.NoError(t, err)
	srv.Log.Debug(resp.String())
}
