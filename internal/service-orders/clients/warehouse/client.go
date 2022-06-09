package warehouse

import (
	"context"
	"fmt"
	"time"

	"github.com/opentracing/opentracing-go"
	"gitlab.ozon.dev/lvjonok/homework-3/pkg/common/api"
	whAPI "gitlab.ozon.dev/lvjonok/homework-3/pkg/srv_warehouse/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type msWarehouseClient struct {
	client whAPI.WarehouseClient

	Timeout time.Duration
}

func New(endpoint string, timeout time.Duration) (*msWarehouseClient, error) {
	conn, err := grpc.Dial(endpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to client, err: <%v>", err)
	}

	return &msWarehouseClient{
		client:  whAPI.NewWarehouseClient(conn),
		Timeout: timeout,
	}, nil
}

func (m *msWarehouseClient) CheckProducts(ctx context.Context, req *whAPI.CheckProductsRequest) (*whAPI.CheckProductsResponse, error) {
	clientCtx, cancel := context.WithTimeout(ctx, m.Timeout)
	defer cancel()

	span, clientCtx := opentracing.StartSpanFromContext(clientCtx, "client CheckProducts")
	defer span.Finish()

	return m.client.CheckProducts(clientCtx, req)
}
func (m *msWarehouseClient) BookProducts(ctx context.Context, req *whAPI.BookProductsRequest) (*whAPI.BookProductsResponse, error) {
	clientCtx, cancel := context.WithTimeout(ctx, m.Timeout)
	defer cancel()

	span, clientCtx := opentracing.StartSpanFromContext(clientCtx, "client BookProducts")
	defer span.Finish()

	return m.client.BookProducts(clientCtx, req)
}
func (m *msWarehouseClient) UnbookProducts(ctx context.Context, req *whAPI.UnbookProductsRequest) (*api.Empty, error) {
	clientCtx, cancel := context.WithTimeout(ctx, m.Timeout)
	defer cancel()

	span, clientCtx := opentracing.StartSpanFromContext(clientCtx, "client UnbookProducts")
	defer span.Finish()

	return m.client.UnbookProducts(clientCtx, req)
}
