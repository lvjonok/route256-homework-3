package repo

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
	"github.com/opentracing/opentracing-go"
	types "gitlab.ozon.dev/lvjonok/homework-3/core/models"
	"gitlab.ozon.dev/lvjonok/homework-3/internal/service-marketplace/models"
)

func (c *Client) UpdateCart(ctx context.Context, cart *models.Cart) (*types.ID, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "database update cart")
	defer span.Finish()

	var b pgx.Batch
	b.Queue(`UPDATE cart SET deleted=TRUE where user_id=$1`, cart.UserID)
	for _, p := range cart.Products {
		b.Queue(`INSERT INTO cart(user_id, product_id, quantity) VALUES ($1, $2, $3)`, cart.UserID, p.ProductID, p.Quantity)
	}

	res := c.pool.SendBatch(ctx, &b)

	var err error
	// var rows pgx.Rows
	for err == nil {
		_, err = res.Query()
	}

	return &cart.UserID, nil
}

func (c *Client) GetCart(ctx context.Context, id *types.ID) (*models.Cart, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "database get cart")
	defer span.Finish()

	const query = `SELECT user_id, product_id, quantity
		FROM cart
		WHERE user_id=$1 AND deleted=FALSE
		ORDER BY created_at DESC;`

	var cart models.Cart
	cart.UserID = *id

	rows, err := c.pool.Query(ctx, query, *id)
	if err != nil {
		return nil, fmt.Errorf("failed to query cart, err: <%v>", err)
	}

	for rows.Next() {
		var p types.ProductUnit
		if err := rows.Scan(&cart.UserID, &p.ProductID, &p.Quantity); err != nil {
			return nil, fmt.Errorf("failed to scan product unit, err: <%v>", err)
		}

		cart.Products = append(cart.Products, p)
	}

	return &cart, nil
}
