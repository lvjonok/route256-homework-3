package service

import (
	"context"

	types "gitlab.ozon.dev/lvjonok/homework-3/core/models"
	"gitlab.ozon.dev/lvjonok/homework-3/internal/service-marketplace/models"
)

type Cache interface {
	GetProduct(context.Context, types.ID) (*models.Product, error)
	UpsertProduct(context.Context, models.Product) error
	GetReviews(context.Context, types.ID) ([]models.Review, error)
	AppendReview(context.Context, models.Review) error
	UpsertReviews(context.Context, types.ID, []models.Review) error
	GetCart(context.Context, types.ID) (*models.Cart, error)
	UpsertCart(context.Context, models.Cart) error
}
