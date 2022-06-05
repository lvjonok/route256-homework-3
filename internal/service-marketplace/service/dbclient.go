package service

import (
	"context"
	"errors"

	types "gitlab.ozon.dev/lvjonok/homework-3/core/models"
	"gitlab.ozon.dev/lvjonok/homework-3/internal/service-marketplace/models"
)

var DBNotFound = errors.New("not found in database")

// DB defines interfaces we want our database client to implement
type DB interface {
	CreateProduct(context.Context, *models.Product) (*types.ID, error)
	GetProduct(context.Context, *types.ID) (*models.Product, error)

	UpdateCart(context.Context, *models.Cart) (*types.ID, error)
	GetCart(context.Context, *types.ID) (*models.Cart, error)

	CreateReview(context.Context, *models.Review) (*types.ID, error)
	GetProductReviews(context.Context, *types.ID) ([]models.Review, error)
}
