package service

import (
	pb "gitlab.ozon.dev/lvjonok/homework-3/pkg/srv_marketplace/api"
	"go.uber.org/zap"
)

type Service struct {
	DB      DB
	Metrics Metrics
	Log     *zap.Logger
	pb.UnimplementedMarketplaceServer
}

func New(db DB, metrics Metrics, logger *zap.Logger) *Service {
	return &Service{
		DB:      db,
		Metrics: metrics,
		Log:     logger,
	}
}
