package repo

import (
	"context"
	"fmt"

	"github.com/opentracing/opentracing-go"
	types "gitlab.ozon.dev/lvjonok/homework-3/core/models"
	"gitlab.ozon.dev/lvjonok/homework-3/internal/service-marketplace/models"
)

func (c *Client) CreateReview(ctx context.Context, review *models.Review) (*types.ID, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "database create review")
	defer span.Finish()

	const query = `INSERT INTO review (product_id, text)
		VALUES ($1, $2)
		RETURNING id;`

	var reviewID types.ID

	if err := c.pool.QueryRow(ctx, query, review.ProductID, review.Text).Scan(&reviewID); err != nil {
		return nil, fmt.Errorf("failed to create review, err: <%v>", err)
	}

	return &reviewID, nil
}

func (c *Client) GetProductReviews(ctx context.Context, id *types.ID) ([]models.Review, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "database get product reviews")
	defer span.Finish()

	const query = `SELECT id, product_id, text
		FROM review
		WHERE product_id = $1;`

	var reviews []models.Review

	// c.pool.

	rows, err := c.pool.Query(ctx, query, *id)
	if err != nil {
		return nil, fmt.Errorf("failed to query reviews, err: <%v>", err)
	}

	for rows.Next() {
		var review models.Review
		if err := rows.Scan(&review.ID, &review.ProductID, &review.Text); err != nil {
			return nil, fmt.Errorf("failed to scan review row, err: <%v>", err)
		}

		reviews = append(reviews, review)
	}

	return reviews, nil
}
