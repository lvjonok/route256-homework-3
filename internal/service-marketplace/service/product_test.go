package service_test

import (
	"context"
	"testing"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	types "gitlab.ozon.dev/lvjonok/homework-3/core/models"
	"gitlab.ozon.dev/lvjonok/homework-3/internal/service-marketplace/models"
	"gitlab.ozon.dev/lvjonok/homework-3/internal/service-marketplace/repo"
	"gitlab.ozon.dev/lvjonok/homework-3/internal/service-marketplace/service"
	service_marketplace "gitlab.ozon.dev/lvjonok/homework-3/pkg/srv_marketplace/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// func Test(t *testing.T) {
// 	mc := minimock.NewController(t)
// 	defer mc.Finish()

// 	mockDB := service.NewDBMock(mc)
// 	mockDB.UpdateCartMock.Return(types.Int2ID(1234), nil)

// 	srv := service.New(mockDB, nil, nil)
// 	ctx := context.Background()

// 	resp, err := srv.UpdateCart(ctx, &service_marketplace.UpdateCartRequest{
// 		ID:       1234,
// 		Products: []*service_marketplace.ProductUnit{},
// 	})
// 	require.NoError(t, err)
// 	require.Equal(t, uint64(1234), resp.ID)
// }

func TestCreateProduct(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	mockDB := service.NewDBMock(mc)
	mockDB.CreateProductMock.Return(types.Int2ID(1234), nil)

	srv := service.New(mockDB, nil, nil)
	ctx := context.Background()

	resp, err := srv.CreateProduct(ctx, &service_marketplace.CreateProductRequest{
		Name: "asdf",
		Desc: "qwer",
	})
	require.NoError(t, err)
	require.Equal(t, uint64(1234), resp.ID)
}

func TestGetProduct(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	mockDB := service.NewDBMock(mc)
	mockDB.GetProductMock.Return(&models.Product{
		ID:   1234,
		Name: "asdf",
		Desc: "qwer",
	}, nil)

	srv := service.New(mockDB, nil, nil)
	ctx := context.Background()

	resp, err := srv.GetProduct(ctx, &service_marketplace.GetProductRequest{
		ID: 1234,
	})
	require.NoError(t, err)
	require.Equal(t, uint64(1234), resp.ID)
	require.Equal(t, "asdf", resp.Name)
	require.Equal(t, "qwer", resp.Desc)
}

func TestGetProductNotFound(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	mockDB := service.NewDBMock(mc)
	mockDB.GetProductMock.Return(nil, repo.ErrNotFound)

	srv := service.New(mockDB, nil, nil)
	ctx := context.Background()

	_, err := srv.GetProduct(ctx, &service_marketplace.GetProductRequest{
		ID: 1234,
	})
	require.Equal(t, codes.NotFound, status.Code(err))
}
