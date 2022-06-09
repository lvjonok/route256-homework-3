package repo_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.ozon.dev/lvjonok/homework-3/core/dbconnector"
	"gitlab.ozon.dev/lvjonok/homework-3/internal/service-warehouse/repo"
)

func Prepare(t *testing.T) (*repo.Client, context.Context) {
	// cfg, err := TODO: add configs

	ctx := context.Background()
	adp, err := dbconnector.New(ctx, "postgresql://root:root@localhost:49263/root")
	require.NoError(t, err)

	_, err = adp.Exec(ctx, "TRUNCATE TABLE entries CASCADE;")
	require.NoError(t, err)

	client := repo.New(adp)
	return client, ctx
}
