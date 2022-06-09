package service

import (
	"context"

	"gitlab.ozon.dev/lvjonok/homework-3/pkg/common/api"
	mpAPI "gitlab.ozon.dev/lvjonok/homework-3/pkg/srv_marketplace/api"
	whAPI "gitlab.ozon.dev/lvjonok/homework-3/pkg/srv_warehouse/api"
)

type MarketplaceClient interface {
	GetCart(context.Context, *mpAPI.GetCartRequest) (*mpAPI.GetCartResponse, error)
}

type WarehouseClient interface {
	CheckProducts(context.Context, *whAPI.CheckProductsRequest) (*whAPI.CheckProductsResponse, error)
	BookProducts(context.Context, *whAPI.BookProductsRequest) (*whAPI.BookProductsResponse, error)
	UnbookProducts(context.Context, *whAPI.UnbookProductsRequest) (*api.Empty, error)
}
