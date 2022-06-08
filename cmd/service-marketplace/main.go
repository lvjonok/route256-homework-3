package main

import (
	"context"
	"net"
	"net/http"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"gitlab.ozon.dev/lvjonok/homework-3/internal/service-marketplace/dbconnector"
	"gitlab.ozon.dev/lvjonok/homework-3/internal/service-marketplace/metrics"
	"gitlab.ozon.dev/lvjonok/homework-3/internal/service-marketplace/mw"
	"gitlab.ozon.dev/lvjonok/homework-3/internal/service-marketplace/repo"
	"gitlab.ozon.dev/lvjonok/homework-3/internal/service-marketplace/service"
	pb "gitlab.ozon.dev/lvjonok/homework-3/pkg/srv_marketplace/api"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
)

func main() {
	log, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	jaegercfg := &config.Configuration{
		ServiceName: "service",
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans: true,
		},
	}

	tracer, closer, err := jaegercfg.NewTracer(config.Logger(jaeger.StdLogger))
	if err != nil {
		log.Sugar().Fatalf("err jaeger new tracer: <%v>", err)
	}
	defer closer.Close()

	opentracing.SetGlobalTracer(
		tracer,
	)

	// initialize metrics handler
	metrics := metrics.New()
	http.Handle("/metrics", promhttp.Handler())
	go http.ListenAndServe("localhost:2112", nil)

	dbconn, err := dbconnector.New(context.Background(), "postgresql://root:root@localhost:5432/root")
	if err != nil {
		log.Sugar().Fatalf("err db connection: <%v>", err)
	}
	newServer := service.New(repo.New(dbconn), metrics, log)

	lis, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Sugar().Fatalf("failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(mw.MetricsInterceptor, mw.SpanInterceptor),
	}

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterMarketplaceServer(grpcServer, newServer)
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Sugar().Fatalf("failed to serve grpc, err: <%v>", err)
	}

	for {
		time.Sleep(time.Second)
	}

}
