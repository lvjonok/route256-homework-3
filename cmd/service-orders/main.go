package main

import (
	"context"
	"net"
	"net/http"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"gitlab.ozon.dev/lvjonok/homework-3/core/dbconnector"
	"gitlab.ozon.dev/lvjonok/homework-3/internal/service-orders/clients/marketplace"
	"gitlab.ozon.dev/lvjonok/homework-3/internal/service-orders/clients/warehouse"
	"gitlab.ozon.dev/lvjonok/homework-3/internal/service-orders/metrics"
	"gitlab.ozon.dev/lvjonok/homework-3/internal/service-orders/mw"
	"gitlab.ozon.dev/lvjonok/homework-3/internal/service-orders/repo"
	"gitlab.ozon.dev/lvjonok/homework-3/internal/service-orders/service"
	pb "gitlab.ozon.dev/lvjonok/homework-3/pkg/srv_orders/api"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	cfg "gitlab.ozon.dev/lvjonok/homework-3/internal/service-orders/config"
)

func main() {
	cfg, err := cfg.New("cmd/service-orders/config.yaml")
	if err != nil {
		panic(err)
	}

	log, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	jaegercfg := &config.Configuration{
		ServiceName: cfg.Service.Name,
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: cfg.Metrics.JaegerURL,
		},
	}

	tracer, closer, err := jaegercfg.NewTracer(config.Logger(jaeger.StdLogger))
	if err != nil {
		log.Sugar().Fatalf("err jaeger new tracer: <%v>", err)
	}
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	// initialize metrics handler
	metrics := metrics.New()
	http.Handle("/metrics", promhttp.Handler())
	go http.ListenAndServe(cfg.Metrics.PrometheusURL, nil)

	dbconn, err := dbconnector.New(context.Background(), cfg.Database.URL)
	if err != nil {
		log.Sugar().Fatalf("err db connection: <%v>", err)
	}

	mpClient, err := marketplace.New(cfg.Clients.Marketplace.URL, time.Duration(cfg.Clients.Marketplace.Timeout)*time.Millisecond)
	if err != nil {
		log.Sugar().Fatalf("err marketplace connection: <%v>", err)
	}

	whClient, err := warehouse.New(cfg.Clients.Warehouse.URL, time.Millisecond*time.Duration(cfg.Clients.Warehouse.Timeout))
	if err != nil {
		log.Sugar().Fatalf("err warehouse connection: <%v>", err)
	}

	newServer := service.New(repo.New(dbconn), metrics, log, mpClient, whClient)

	lis, err := net.Listen("tcp", cfg.Server.URL)
	if err != nil {
		log.Sugar().Fatalf("failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(mw.MetricsInterceptor, mw.SpanInterceptor),
	}

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterOrdersServer(grpcServer, newServer)
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Sugar().Fatalf("failed to serve grpc, err: <%v>", err)
	}

	for {
		time.Sleep(time.Second)
	}

}
