package mw

import (
	"context"
	"reflect"
	"runtime"

	"github.com/opentracing/opentracing-go"
	"gitlab.ozon.dev/lvjonok/homework-3/internal/service-warehouse/service"
	"google.golang.org/grpc"
)

func MetricsInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	srv := info.Server.(*service.Service)
	srv.Metrics.RequestsInc()

	resp, err := handler(ctx, req)
	if err != nil {
		srv.Metrics.RequestsErrorsInc()
		// TODO: change to error method
		srv.Log.Sugar().Debug(info.FullMethod, req, err)
	} else {
		srv.Log.Sugar().Debug(info.FullMethod, req)
	}
	return resp, err
}

func SpanInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// obtain handler name to start spanning
	handlerFunc := runtime.FuncForPC(reflect.ValueOf(handler).Pointer())

	span, ctx := opentracing.StartSpanFromContext(ctx, handlerFunc.Name())
	defer span.Finish()
	return handler(ctx, req)
}
