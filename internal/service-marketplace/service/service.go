package service

import (
	pb "gitlab.ozon.dev/lvjonok/homework-3/pkg/srv_marketplace/api"
)

type Service struct {
	DB      DB
	Metrics Metrics
	pb.UnimplementedMarketplaceServer
}

func New(db DB, metrics Metrics) *Service {
	return &Service{
		DB:      db,
		Metrics: metrics,
	}
}
