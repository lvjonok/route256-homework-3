package service

import (
	"context"

	types "gitlab.ozon.dev/lvjonok/homework-3/core/models"
	"gitlab.ozon.dev/lvjonok/homework-3/internal/service-warehouse/models"
)

type DB interface {
	RegisterProduct(context.Context, *models.Entry) error
	CheckProducts(context.Context, []types.ID) ([]types.ProductUnit, error)
	BookProducts(context.Context, []types.ProductUnit) ([]types.ID, error)
	UnbookProducts(context.Context, []types.ID) error
}
