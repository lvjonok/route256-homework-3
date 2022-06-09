package repo

import (
	"errors"

	"github.com/jackc/pgx/v4/pgxpool"
)

var ErrNotFound = errors.New("not found")
var ErrNotEnough = errors.New("not enough")

type Client struct {
	pool *pgxpool.Pool
}

func New(pool *pgxpool.Pool) *Client {
	return &Client{pool: pool}
}
