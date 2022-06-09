package service

import (
	pb "gitlab.ozon.dev/lvjonok/homework-3/pkg/srv_orders/api"
	"go.uber.org/zap"
)

type Service struct {
	DB      DB
	Metrics Metrics
	Log     *zap.Logger
	pb.UnimplementedOrdersServer

	mpClient MarketplaceClient
	whClient WarehouseClient
}

func New(db DB, metrics Metrics, logger *zap.Logger, mpClient MarketplaceClient, whClient WarehouseClient) *Service {
	return &Service{
		DB:       db,
		Metrics:  metrics,
		Log:      logger,
		mpClient: mpClient,
		whClient: whClient,
	}
}
