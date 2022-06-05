package main

import (
	"log"
	"net"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"gitlab.ozon.dev/lvjonok/homework-3/internal/service-marketplace/metrics"
	"gitlab.ozon.dev/lvjonok/homework-3/internal/service-marketplace/mw"
	"gitlab.ozon.dev/lvjonok/homework-3/internal/service-marketplace/service"
	pb "gitlab.ozon.dev/lvjonok/homework-3/pkg/srv_marketplace/api"
	"google.golang.org/grpc"
)

var counter = promauto.NewCounter(prometheus.CounterOpts{
	Name: "orders_processed",
})

func main() {
	// initialize metrics handler
	metrics := metrics.New()
	http.Handle("/metrics", promhttp.Handler())
	go http.ListenAndServe("localhost:2112", nil)

	// cfg, err := config.New("config.yaml")
	// if err != nil {
	// 	panic(err)
	// }
	go func() {
		for {
			counter.Inc()
			// simulate some processing function
			time.Sleep(time.Second)
		}
	}()

	newServer := service.New(nil, metrics)

	lis, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(mw.MetricsInterceptor))
	pb.RegisterMarketplaceServer(grpcServer, newServer)
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve grpc, err: <%v>", err)
	}

	for {
		time.Sleep(time.Second)
	}

}
