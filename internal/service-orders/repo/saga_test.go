package repo_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	types "gitlab.ozon.dev/lvjonok/homework-3/core/models"
	"gitlab.ozon.dev/lvjonok/homework-3/internal/service-orders/models"
)

func TestAddRetry(t *testing.T) {
	client, ctx := Prepare(t)

	err := client.AddRetry(ctx, *types.Int2ID(10), models.Checked)
	require.NoError(t, err)
}

func TestGetProcessingOrders(t *testing.T) {
	client, ctx := Prepare(t)

	// create test order
	order := models.Order{
		UserID:     0,
		SagaStatus: models.Created,
	}

	orderid, err := client.CreateOrder(ctx, &order)
	require.NoError(t, err)

	err = client.AddRetry(ctx, *orderid, models.Checked)
	require.NoError(t, err)
	err = client.AddRetry(ctx, *orderid, models.Checked)
	require.NoError(t, err)

	// after two retries we should get this order
	res, err := client.GetProcessingOrders(ctx, 5)
	require.NoError(t, err)
	require.Len(t, res, 1)
	require.Equal(t, *orderid, res[0])

	// if after attempt it became finished, we should not get it
	err = client.UpdateOrderSagaStatus(ctx, &models.Order{
		OrderID:    *orderid,
		SagaStatus: models.Booked,
	})
	require.NoError(t, err)

	res, err = client.GetProcessingOrders(ctx, 5)
	require.NoError(t, err)
	require.Len(t, res, 0)

	// after 5 retries, we should just accept that it failed
	err = client.UpdateOrderSagaStatus(ctx, &models.Order{
		OrderID:    *orderid,
		SagaStatus: models.Checked,
	})
	require.NoError(t, err)
	res, err = client.GetProcessingOrders(ctx, 5)
	require.NoError(t, err)
	require.Len(t, res, 1)

	require.NoError(t, err)
	err = client.AddRetry(ctx, *orderid, models.Checked)
	require.NoError(t, err)
	err = client.AddRetry(ctx, *orderid, models.Checked)
	require.NoError(t, err)
	err = client.AddRetry(ctx, *orderid, models.Checked)
	require.NoError(t, err)

	res, err = client.GetProcessingOrders(ctx, 5)
	require.NoError(t, err)
	require.Len(t, res, 0)
}
