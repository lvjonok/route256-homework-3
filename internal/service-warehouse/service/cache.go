package service

import (
	"context"

	types "gitlab.ozon.dev/lvjonok/homework-3/core/models"
)

type Cache interface {
	// GetProducts tries to get all products by ids from cache
	GetProducts(context.Context, []types.ID) ([]types.ProductUnit, error)
	// UpsertProducts updates values for products in cache
	UpsertProducts(context.Context, []types.ProductUnit) error
	// DeleteProducts removes ids of products from cache, so we would refresh them
	DeleteProducts(context.Context, []types.ID) error
}
