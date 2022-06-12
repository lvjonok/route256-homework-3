package cache

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/opentracing/opentracing-go"
	"gitlab.ozon.dev/lvjonok/homework-3/core/cacheconnector"
	types "gitlab.ozon.dev/lvjonok/homework-3/core/models"
	"gitlab.ozon.dev/lvjonok/homework-3/internal/service-marketplace/models"
)

func (c *cache) GetCart(ctx context.Context, id types.ID) (*models.Cart, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "CacheGetCart")
	defer span.Finish()

	item, err := c.client.Get(fmt.Sprintf("cart%v", id))
	if errors.Is(err, memcache.ErrCacheMiss) {
		return nil, cacheconnector.ErrCacheMiss
	} else if err != nil {
		return nil, fmt.Errorf("failed to get cart: <%v>", err)
	}

	var cart models.Cart
	if err := json.Unmarshal(item.Value, &cart); err != nil {
		return nil, fmt.Errorf("failed to unmarshall cart, err: <%v>", err)
	}

	return &cart, nil
}

func (c *cache) UpsertCart(ctx context.Context, cart models.Cart) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "CacheUpsertCart")
	defer span.Finish()

	key := fmt.Sprintf("cart%v", cart.UserID)
	value, err := json.Marshal(cart)
	if err != nil {
		return fmt.Errorf("failed to marshall cart, err: <%v>", err)
	}

	return c.client.Set(&memcache.Item{
		Key:   key,
		Value: value,
	})
}
