package repo

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v4"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	types "gitlab.ozon.dev/lvjonok/homework-3/core/models"
	"gitlab.ozon.dev/lvjonok/homework-3/internal/service-warehouse/models"
)

func (c *Client) RegisterProduct(ctx context.Context, el *models.Entry) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "db RegisterProduct")
	defer span.Finish()

	const query = `INSERT INTO entries(product_id, quantity)
		VALUES ($1, $2)
		RETURNING id;`

	var id types.ID
	if err := c.pool.QueryRow(ctx, query, el.ProductID, el.Quantity).Scan(&id); err != nil {
		return fmt.Errorf("failed to register product, err: <%v>", err)
	}

	return nil
}

func (c *Client) CheckProducts(ctx context.Context, el []types.ID) ([]types.ProductUnit, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "db CheckProducts")
	defer span.Finish()

	var units []types.ProductUnit

	indices := []string{}
	for _, e := range el {
		indices = append(indices, strconv.Itoa(int(e)))
	}

	query := `SELECT product_id, SUM(quantity)
		FROM entries
		WHERE deleted = FALSE
		AND product_id IN (` + strings.Join(indices, ",") + `)
		GROUP BY product_id;`

	rows, err := c.pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to query products, err: <%v>", err)
	}

	for rows.Next() {
		var unit types.ProductUnit
		if err := rows.Scan(&unit.ProductID, &unit.Quantity); err != nil {
			return nil, fmt.Errorf("failed to scan unit, err: <%v>", err)
		}
		units = append(units, unit)
	}

	return units, nil
}

func (c *Client) BookProducts(ctx context.Context, el []types.ProductUnit) ([]types.ID, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "db BookProducts")
	defer span.Finish()

	var ids []types.ID

	tx, err := c.pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to start transaction, err: <%v>", err)
	}
	defer tx.Rollback(ctx)

	const query = `INSERT INTO entries (product_id, quantity)
		VALUES ($1, $2)
		RETURNING id;`

	const query2 = `SELECT SUM(quantity)
		FROM entries
		WHERE deleted = FALSE
		AND product_id = $1
		GROUP BY product_id;`

	for _, e := range el {
		var q int
		if err := tx.QueryRow(ctx, query2, e.ProductID).Scan(&q); err != nil {
			return nil, fmt.Errorf("failed to scan current quantity of product, err: <%v>", err)
		}

		// check that we have enough
		if q < e.Quantity {
			return nil, errors.Wrapf(ErrNotEnough, "product: <%v>, want: %v, have: %v", e.ProductID, e.Quantity, q)
		}

		var id types.ID
		if err := tx.QueryRow(ctx, query, e.ProductID, -e.Quantity).Scan(&id); err != nil {
			return nil, fmt.Errorf("failed to scan entry, err: <%v>", err)
		}
		ids = append(ids, id)
	}

	return ids, tx.Commit(ctx)
}

func (c *Client) UnbookProducts(ctx context.Context, el []types.ID) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "db UnbookProducts")
	defer span.Finish()

	indices := []string{}
	for _, e := range el {
		indices = append(indices, strconv.Itoa(int(e)))
	}

	query := `UPDATE entries
		SET deleted=TRUE
		WHERE id IN (` + strings.Join(indices, ",") + `);`

	r, err := c.pool.Exec(ctx, query)
	if err != nil || r.RowsAffected() == 0 {
		return fmt.Errorf("failed to exec update unbook products, err: <%v>", err)
	}

	return nil
}
