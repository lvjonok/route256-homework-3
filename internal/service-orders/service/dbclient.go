package service

import (
	"context"
	"errors"

	types "gitlab.ozon.dev/lvjonok/homework-3/core/models"
	"gitlab.ozon.dev/lvjonok/homework-3/internal/service-orders/models"
)

var ErrNotFound = errors.New("not found in database")

// DB defines interfaces we want our database client to implement
type DB interface {
	CreateOrder(context.Context, *models.Order) (*types.ID, error)
	CheckStatus(context.Context, *types.ID) (string, error)
	UpdateStatus(context.Context, *models.Order) error
}
