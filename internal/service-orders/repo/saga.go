package repo

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
	"github.com/opentracing/opentracing-go"
	types "gitlab.ozon.dev/lvjonok/homework-3/core/models"
	"gitlab.ozon.dev/lvjonok/homework-3/internal/service-orders/models"
)

func (c *Client) UpdateOrderSagaStatus(ctx context.Context, order *models.Order) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "db UpdateOrder")
	defer span.Finish()

	const query = `UPDATE user_orders SET saga_status=$2 WHERE order_id=$1 RETURNING 1`

	var x int
	if err := c.pool.QueryRow(ctx, query, order.OrderID, order.SagaStatus).Scan(&x); err != nil {
		if err == pgx.ErrNoRows {
			return ErrNotFound
		}
		return fmt.Errorf("failed to update order saga status, err: <%v>", err)
	}

	return nil
}

func (c *Client) GetProcessingOrders(ctx context.Context, retries int) ([]types.ID, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "db GetProcessingOrders")
	defer span.Finish()

	const query = `SELECT user_orders.order_id
		FROM user_orders
				JOIN (SELECT order_id, COUNT(*) AS counter
					FROM retries
					GROUP BY order_id) AS retries
					ON user_orders.order_id = retries.order_id
		WHERE user_orders.saga_status != $1
		AND retries.counter < $2`

	var ids []types.ID
	rows, err := c.pool.Query(ctx, query, models.Booked, retries)
	if err != nil {
		return nil, fmt.Errorf("failed to query processing orders, err: <%v>", err)
	}

	for rows.Next() {
		var id types.ID
		if err := rows.Scan(&id); err != nil {
			return nil, fmt.Errorf("failed to query id ")
		}
		ids = append(ids, id)
	}

	return ids, nil
}

func (c *Client) AddRetry(ctx context.Context, order_id types.ID, saga models.OrderStatus) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "db AddRetry")
	defer span.Finish()

	const query = `INSERT INTO retries(order_id, last_status) VALUES ($1, $2) RETURNING 1;`

	var x int
	if err := c.pool.QueryRow(ctx, query, order_id, saga).Scan(&x); err != nil {
		return fmt.Errorf("failed to add retry, err: <%v>", err)
	}

	return nil
}
