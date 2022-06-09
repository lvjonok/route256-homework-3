package service

import (
	"context"

	mpAPI "gitlab.ozon.dev/lvjonok/homework-3/pkg/srv_marketplace/api"
)

type MarketplaceClient interface {
	GetProduct(context.Context, *mpAPI.GetProductRequest) (*mpAPI.GetProductResponse, error)
}
