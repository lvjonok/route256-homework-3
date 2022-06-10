package repo_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.ozon.dev/lvjonok/homework-3/core/dbconnector"
	"gitlab.ozon.dev/lvjonok/homework-3/internal/service-marketplace/config"
	"gitlab.ozon.dev/lvjonok/homework-3/internal/service-marketplace/repo"
)

func Prepare(t *testing.T) (*repo.Client, context.Context) {
	cfg, err := config.New("../../../cmd/service-marketplace/config.yaml")
	require.NoError(t, err)

	ctx := context.Background()
	adp, err := dbconnector.New(ctx, cfg.Database.URL)
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
