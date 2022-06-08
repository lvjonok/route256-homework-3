package repo_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	types "gitlab.ozon.dev/lvjonok/homework-3/core/models"
	"gitlab.ozon.dev/lvjonok/homework-3/internal/service-marketplace/models"
	"gitlab.ozon.dev/lvjonok/homework-3/internal/service-marketplace/repo"
)

func TestCreateProduct(t *testing.T) {
	client, ctx := Prepare(t)

	_, err := client.CreateProduct(ctx, &models.Product{
		Name: "name",
		Desc: "desc",
	})
	require.NoError(t, err)
}

func TestGetProduct(t *testing.T) {
	client, ctx := Prepare(t)

	id, err := client.CreateProduct(ctx, &models.Product{
		Name: "name",
		Desc: "desc",
	})
	require.NoError(t, err)

	p, err := client.GetProduct(ctx, id)
	require.NoError(t, err)

	assert.Equal(t, "name", p.Name)
	assert.Equal(t, "desc", p.Desc)
}

func TestGetProductNotFound(t *testing.T) {
	client, ctx := Prepare(t)

	id := types.ID(1)

	_, err := client.GetProduct(ctx, &id)
	require.Equal(t, repo.ErrNotFound, err)
}
