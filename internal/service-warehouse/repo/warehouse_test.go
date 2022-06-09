package repo_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	types "gitlab.ozon.dev/lvjonok/homework-3/core/models"
	"gitlab.ozon.dev/lvjonok/homework-3/internal/service-warehouse/models"
	"gitlab.ozon.dev/lvjonok/homework-3/internal/service-warehouse/repo"
)

func TestRegisterProduct(t *testing.T) {
	client, ctx := Prepare(t)
	err := client.RegisterProduct(ctx, &models.Entry{
		ProductID: 1234,
		Quantity:  10,
	})
	require.NoError(t, err)
}

func TestCheckProducts(t *testing.T) {
	client, ctx := Prepare(t)
	err := client.RegisterProduct(ctx, &models.Entry{
		ProductID: 1,
		Quantity:  10,
	})
	require.NoError(t, err)
	err = client.RegisterProduct(ctx, &models.Entry{
		ProductID: 2,
		Quantity:  20,
	})
	require.NoError(t, err)
	err = client.RegisterProduct(ctx, &models.Entry{
		ProductID: 3,
		Quantity:  50,
	})
	require.NoError(t, err)

	units, err := client.CheckProducts(ctx, []types.ID{
		*types.Int2ID(1), *types.Int2ID(2),
	})
	require.NoError(t, err)

	require.Len(t, units, 2)
	require.Equal(t, 10, units[0].Quantity)
	require.Equal(t, 20, units[1].Quantity)
}

func TestBookProducs(t *testing.T) {
	client, ctx := Prepare(t)
	err := client.RegisterProduct(ctx, &models.Entry{
		ProductID: 1,
		Quantity:  10,
	})
	require.NoError(t, err)
	err = client.RegisterProduct(ctx, &models.Entry{
		ProductID: 2,
		Quantity:  20,
	})
	require.NoError(t, err)
	err = client.RegisterProduct(ctx, &models.Entry{
		ProductID: 3,
		Quantity:  50,
	})
	require.NoError(t, err)

	_, err = client.BookProducts(ctx, []types.ProductUnit{
		{
			ProductID: 1,
			Quantity:  10,
		},
		{
			ProductID: 2,
			Quantity:  20,
		},
	})
	require.NoError(t, err)
}

func TestBookProductsNotEnough(t *testing.T) {
	client, ctx := Prepare(t)
	err := client.RegisterProduct(ctx, &models.Entry{
		ProductID: 1,
		Quantity:  1,
	})
	require.NoError(t, err)
	err = client.RegisterProduct(ctx, &models.Entry{
		ProductID: 2,
		Quantity:  1,
	})
	require.NoError(t, err)

	_, err = client.BookProducts(ctx, []types.ProductUnit{
		{
			ProductID: 1,
			Quantity:  10,
		},
		{
			ProductID: 2,
			Quantity:  20,
		},
	})
	require.ErrorIs(t, err, repo.ErrNotEnough)
}

func TestUnbookProducts(t *testing.T) {
	client, ctx := Prepare(t)
	err := client.RegisterProduct(ctx, &models.Entry{
		ProductID: 1,
		Quantity:  10,
	})
	require.NoError(t, err)
	err = client.RegisterProduct(ctx, &models.Entry{
		ProductID: 2,
		Quantity:  20,
	})
	require.NoError(t, err)

	ids, err := client.BookProducts(ctx, []types.ProductUnit{
		{
			ProductID: 1,
			Quantity:  10,
		},
		{
			ProductID: 2,
			Quantity:  20,
		},
	})
	require.NoError(t, err)

	err = client.UnbookProducts(ctx, ids)
	require.NoError(t, err)

	units, err := client.CheckProducts(ctx, []types.ID{1, 2})
	require.NoError(t, err)

	require.Len(t, units, 2)
	require.Equal(t, 10, units[0].Quantity)
	require.Equal(t, 20, units[1].Quantity)
}
