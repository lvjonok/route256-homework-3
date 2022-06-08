package repo_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.ozon.dev/lvjonok/homework-3/internal/service-marketplace/models"
)

func TestUpdateCart(t *testing.T) {
	client, ctx := Prepare(t)

	p1 := models.Product{
		Name: "p1",
		Desc: "d1",
	}
	p2 := models.Product{
		Name: "p2",
		Desc: "d2",
	}

	p1id, err := client.CreateProduct(ctx, &p1)
	require.NoError(t, err)
	p1.ID = *p1id

	p2id, err := client.CreateProduct(ctx, &p2)
	require.NoError(t, err)
	p2.ID = *p2id

	_, err = client.UpdateCart(ctx, &models.Cart{
		UserID: 1234,
		Products: []models.ProductUnit{
			{
				ID:       *p1id,
				Quantity: 2,
			},
			{
				ID:       *p2id,
				Quantity: 4,
			},
		},
	})
	require.NoError(t, err)
}

func TestGetCart(t *testing.T) {
	client, ctx := Prepare(t)

	p1 := models.Product{
		Name: "p1",
		Desc: "d1",
	}
	p2 := models.Product{
		Name: "p2",
		Desc: "d2",
	}

	p1id, err := client.CreateProduct(ctx, &p1)
	require.NoError(t, err)
	p1.ID = *p1id

	p2id, err := client.CreateProduct(ctx, &p2)
	require.NoError(t, err)
	p2.ID = *p2id

	cartID, err := client.UpdateCart(ctx, &models.Cart{
		UserID: 1234,
		Products: []models.ProductUnit{
			{
				ID:       *p1id,
				Quantity: 2,
			},
			{
				ID:       *p2id,
				Quantity: 4,
			},
		},
	})
	require.NoError(t, err)

	res, err := client.GetCart(ctx, cartID)
	require.NoError(t, err)
	require.ElementsMatch(t, res.Products, []models.ProductUnit{
		{
			ID:       *p1id,
			Quantity: 2,
		},
		{
			ID:       *p2id,
			Quantity: 4,
		},
	})
}
