package repo_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.ozon.dev/lvjonok/homework-3/internal/service-marketplace/dbconnector"
	"gitlab.ozon.dev/lvjonok/homework-3/internal/service-marketplace/repo"
)

func Prepare(t *testing.T) (*repo.Client, context.Context) {
	// cfg, err := TODO: add configs

	ctx := context.Background()
	adp, err := dbconnector.New(ctx, "postgresql://root:root@localhost:5432/root")
	require.NoError(t, err)

	_, err = adp.Exec(ctx, "TRUNCATE TABLE product CASCADE;")
	require.NoError(t, err)
	_, err = adp.Exec(ctx, "TRUNCATE TABLE review CASCADE;")
	require.NoError(t, err)
	_, err = adp.Exec(ctx, "TRUNCATE TABLE cart CASCADE;")
	require.NoError(t, err)

	client := repo.New(adp)
	return client, ctx
}
