package repo_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.ozon.dev/lvjonok/homework-3/internal/service-marketplace/models"
)

func TestCreateReview(t *testing.T) {
	client, ctx := Prepare(t)

	pid, err := client.CreateProduct(ctx, &models.Product{
		Name: "name",
		Desc: "desc",
	})
	require.NoError(t, err)

	_, err = client.CreateReview(ctx, &models.Review{
		Text:      "text",
		ProductID: *pid,
	})
	require.NoError(t, err)
}

func TestGetProductReviews(t *testing.T) {
	client, ctx := Prepare(t)

	pid1, err := client.CreateProduct(ctx, &models.Product{
		Name: "name",
		Desc: "desc",
	})
	require.NoError(t, err)

	r1, err := client.CreateReview(ctx, &models.Review{
		Text:      "text1",
		ProductID: *pid1,
	})
	require.NoError(t, err)

	r2, err := client.CreateReview(ctx, &models.Review{
		Text:      "text2",
		ProductID: *pid1,
	})
	require.NoError(t, err)

	reviews, err := client.GetProductReviews(ctx, pid1)
	require.NoError(t, err)

	require.Len(t, reviews, 2)

	require.Equal(t, reviews[0].ID, *r1)
	require.Equal(t, reviews[0].ProductID, *pid1)
	require.Equal(t, reviews[0].Text, "text1")

	require.Equal(t, reviews[1].ID, *r2)
	require.Equal(t, reviews[1].ProductID, *pid1)
	require.Equal(t, reviews[1].Text, "text2")
}
