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
)

func TestUpdateCart(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	mockDB := service.NewDBMock(mc)
	mockDB.UpdateCartMock.Return(types.Int2ID(1234), nil)

	srv := service.New(mockDB, nil, nil)
	ctx := context.Background()

	resp, err := srv.UpdateCart(ctx, &service_marketplace.UpdateCartRequest{
		ID:       1234,
		Products: []*service_marketplace.ProductUnit{},
	})
	require.NoError(t, err)
	require.Equal(t, uint64(1234), resp.ID)
}

func TestGetCart(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	cart := models.Cart{
		UserID: 1234,
		Products: []models.ProductUnit{
			{
				ID:       1,
				Quantity: 2,
			},
			{
				ID:       2,
				Quantity: 3,
			},
		},
	}

	mockDB := service.NewDBMock(mc)
	mockDB.GetCartMock.Return(&cart, nil)

	srv := service.New(mockDB, nil, nil)
	ctx := context.Background()

	resp, err := srv.GetCart(ctx, &service_marketplace.GetCartRequest{
		ID: 1234,
	})
	require.NoError(t, err)

	require.Len(t, resp.Products, 2)

	require.Equal(t, uint64(1), resp.Products[0].ProductID)
	require.Equal(t, uint64(2), resp.Products[0].Quantity)

	require.Equal(t, uint64(2), resp.Products[1].ProductID)
	require.Equal(t, uint64(3), resp.Products[1].Quantity)
}
