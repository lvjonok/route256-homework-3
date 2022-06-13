package service

import (
	pb "gitlab.ozon.dev/lvjonok/homework-3/pkg/srv_warehouse/api"
	"go.uber.org/zap"
)

type Service struct {
	DB      DB
	Cache   Cache
	Metrics Metrics
	Log     *zap.Logger
	pb.UnimplementedWarehouseServer

	mpClient MarketplaceClient
}

func New(db DB, cache Cache, metrics Metrics, logger *zap.Logger, mpClient MarketplaceClient) *Service {
	return &Service{
		DB:       db,
		Cache:    cache,
		Metrics:  metrics,
		Log:      logger,
		mpClient: mpClient,
	}
}
