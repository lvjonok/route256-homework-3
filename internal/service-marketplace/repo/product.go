package repo

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4"
	"github.com/opentracing/opentracing-go"
	types "gitlab.ozon.dev/lvjonok/homework-3/core/models"
	"gitlab.ozon.dev/lvjonok/homework-3/internal/service-marketplace/models"
)

func (c *Client) CreateProduct(ctx context.Context, p *models.Product) (*types.ID, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "database create product")
	defer span.Finish()

	const query = `INSERT INTO product(name, "desc")
		VALUES ($1, $2)
		RETURNING id;`

	var id types.ID

	if err := c.pool.QueryRow(ctx, query, p.Name, p.Desc).Scan(&id); err != nil {
		return nil, fmt.Errorf("failed to create product, err: <%v>", err)
	}

	return &id, nil
}

func (c *Client) GetProduct(ctx context.Context, id *types.ID) (*models.Product, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "database get product")
	defer span.Finish()

	const query = `SELECT id, name, "desc"
		FROM product
		WHERE id = $1`

	var p models.Product

	err := c.pool.QueryRow(ctx, query, *id).Scan(&p.ID, &p.Name, &p.Desc)
	if err != nil {
		log.Printf("err <%v> <%v> = %v, %v, %v", err.Error(), pgx.ErrNoRows.Error(), err.Error() == pgx.ErrNoRows.Error(), err == pgx.ErrNoRows, err == sql.ErrNoRows)
		if err == pgx.ErrNoRows {
			return nil, ErrNotFound
		}

		return nil, fmt.Errorf("failed to query product, err: <%v>", err)
	}

	return &p, nil
}
