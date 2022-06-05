package mw

import (
	"context"
	"log"

	"gitlab.ozon.dev/lvjonok/homework-3/internal/service-marketplace/service"
	"google.golang.org/grpc"
)

func LogInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	resp, err := handler(ctx, req)
	if err != nil {
		log.Println(info.FullMethod, req, err)
	} else {
		log.Println(info.FullMethod, req)
	}
	return resp, err
}

func MetricsInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	srv := info.Server.(*service.Service)
	srv.Metrics.RequestsInc()

	resp, err := handler(ctx, req)
	if err != nil {
		srv.Metrics.RequestErrorsInc()
		log.Println(info.FullMethod, req, err)
	} else {
		log.Println(info.FullMethod, req)
	}
	return resp, err
}
