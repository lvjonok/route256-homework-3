package repo_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	types "gitlab.ozon.dev/lvjonok/homework-3/core/models"
	"gitlab.ozon.dev/lvjonok/homework-3/internal/service-orders/models"
)

func TestCreateOrder(t *testing.T) {
	client, ctx := Prepare(t)

	_, err := client.CreateOrder(ctx, &models.Order{
		UserID: 1234,
		Products: []types.ProductUnit{
			{
				ProductID: 1,
				Quantity:  1,
			},
			{
				ProductID: 2,
				Quantity:  2,
			},
		},
	})
	require.NoError(t, err)

}

func TestCheckStatus(t *testing.T) {
	client, ctx := Prepare(t)

	id, err := client.CreateOrder(ctx, &models.Order{
		UserID: 1234,
		Products: []types.ProductUnit{
			{
				ProductID: 1,
				Quantity:  1,
			},
			{
				ProductID: 2,
				Quantity:  2,
			},
		},
		Status: "some really cool status",
	})
	require.NoError(t, err)

	status, err := client.CheckStatus(ctx, id)
	require.NoError(t, err)

	require.Equal(t, "some really cool status", status)
}

func TestUpdateStatus(t *testing.T) {
	client, ctx := Prepare(t)

	id, err := client.CreateOrder(ctx, &models.Order{
		UserID: 1234,
		Products: []types.ProductUnit{
			{
				ProductID: 1,
				Quantity:  1,
			},
			{
				ProductID: 2,
				Quantity:  2,
			},
		},
		Status: "some really cool status",
	})
	require.NoError(t, err)

	err = client.UpdateStatus(ctx, &models.Order{
		OrderID: *id,
		Status:  "cool",
	})
	require.NoError(t, err)

	status, err := client.CheckStatus(ctx, id)
	require.NoError(t, err)

	require.Equal(t, "cool", status)
}

func TestGetOrder(t *testing.T) {
	client, ctx := Prepare(t)

	units := []types.ProductUnit{
		{
			ProductID: 1,
			Quantity:  1,
		},
		{
			ProductID: 2,
			Quantity:  2,
		},
	}

	id, err := client.CreateOrder(ctx, &models.Order{
		UserID:   1234,
		Products: units,
	})
	require.NoError(t, err)

	res, err := client.GetOrder(ctx, id)
	require.NoError(t, err)

	require.ElementsMatch(t, units, res.Products)
}
