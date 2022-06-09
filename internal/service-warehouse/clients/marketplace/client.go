package marketplace

import (
	"context"
	"fmt"
	"time"

	"github.com/opentracing/opentracing-go"
	mpAPI "gitlab.ozon.dev/lvjonok/homework-3/pkg/srv_marketplace/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type msMarketplaceClient struct {
	client mpAPI.MarketplaceClient

	Timeout time.Duration
}

func New(endpoint string, timeout time.Duration) (*msMarketplaceClient, error) {
	conn, err := grpc.Dial(endpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to client, err: <%v>", err)
	}

	return &msMarketplaceClient{
		client:  mpAPI.NewMarketplaceClient(conn),
		Timeout: timeout,
	}, nil
}

func (m *msMarketplaceClient) GetProduct(ctx context.Context, req *mpAPI.GetProductRequest) (*mpAPI.GetProductResponse, error) {
	clientCtx, cancel := context.WithTimeout(ctx, m.Timeout)
	defer cancel()

	span, clientCtx := opentracing.StartSpanFromContext(clientCtx, "client GetProduct")
	defer span.Finish()

	return m.client.GetProduct(clientCtx, req)
}
