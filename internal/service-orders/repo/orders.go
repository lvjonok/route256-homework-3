package repo

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
	"github.com/opentracing/opentracing-go"
	types "gitlab.ozon.dev/lvjonok/homework-3/core/models"
	"gitlab.ozon.dev/lvjonok/homework-3/internal/service-orders/models"
)

func (c *Client) CreateOrder(ctx context.Context, o *models.Order) (*types.ID, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "db CreateOrder")
	defer span.Finish()

	tx, err := c.pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to start transaction, err: <%v>", err)
	}
	defer tx.Rollback(ctx)

	const query1 = `INSERT INTO user_orders(user_id, status)
		VALUES ($1, $2)
		RETURNING order_id;`

	var orderID types.ID

	if err := tx.QueryRow(ctx, query1, o.UserID, o.Status).Scan(&orderID); err != nil {
		return nil, fmt.Errorf("failed to query id of inserted order, err: <%v>", err)
	}

	const query2 = `INSERT INTO order_items(order_id, product_id, quantity)
		VALUES ($1, $2, $3)
		RETURNING 1;`

	for _, item := range o.Products {
		var x int
		if err := tx.QueryRow(ctx, query2, orderID, item.ProductID, item.Quantity).Scan(&x); err != nil {
			return nil, fmt.Errorf("failed to query 1 of inserted order-product, err: <%v>", err)
		}
	}

	return &orderID, tx.Commit(ctx)
}

func (c *Client) CheckStatus(ctx context.Context, id *types.ID) (string, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "db CheckStatus")
	defer span.Finish()

	const query = `SELECT status
		FROM user_orders
		WHERE order_id = $1;`

	var status string
	if err := c.pool.QueryRow(ctx, query, id).Scan(&status); err != nil {
		if err == pgx.ErrNoRows {
			return "", ErrNotFound
		}
		return "", fmt.Errorf("failed to query status of order, err: <%v>", err)
	}

	return status, nil
}

func (c *Client) UpdateStatus(ctx context.Context, m *models.Order) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "db UpdateStatus")
	defer span.Finish()

	const query = `UPDATE user_orders
		SET status = $1
		WHERE order_id = $2 RETURNING 1;`

	var x int
	err := c.pool.QueryRow(ctx, query, m.Status, m.OrderID).Scan(&x)
	if err != nil {
		if err == pgx.ErrNoRows {
			return ErrNotFound
		}

		return fmt.Errorf("failed updating status, err: <%v>", err)
	}

	return nil
}
