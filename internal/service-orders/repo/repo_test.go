package repo_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.ozon.dev/lvjonok/homework-3/core/dbconnector"
	"gitlab.ozon.dev/lvjonok/homework-3/internal/service-orders/repo"
)

func Prepare(t *testing.T) (*repo.Client, context.Context) {
	// cfg, err := TODO: add configs

	ctx := context.Background()
	adp, err := dbconnector.New(ctx, "postgresql://root:root@localhost:49262/root")
	require.NoError(t, err)

	_, err = adp.Exec(ctx, "TRUNCATE TABLE user_orders CASCADE;")
	require.NoError(t, err)
	_, err = adp.Exec(ctx, "TRUNCATE TABLE order_items CASCADE;")
	require.NoError(t, err)

	client := repo.New(adp)
	return client, ctx
}
