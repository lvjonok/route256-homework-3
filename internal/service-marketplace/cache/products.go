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

// GetProduct tries to get product model from cache by id, returns ErrCacheMiss if misses
func (c *cache) GetProduct(ctx context.Context, id types.ID) (*models.Product, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "CacheGetProduct")
	defer span.Finish()

	item, err := c.client.Get(fmt.Sprintf("product%v", id))
	if errors.Is(err, memcache.ErrCacheMiss) {
		return nil, cacheconnector.ErrCacheMiss
	} else if err != nil {
		return nil, fmt.Errorf("failed to get from cache, err: <%v>", err)
	}

	var product models.Product
	if err := json.Unmarshal(item.Value, &product); err != nil {
		return nil, fmt.Errorf("failed to unmarshall product, err: <%v>", err)
	}

	return &product, nil
}

// UpsertProduct updates product in the cache, comparing with the previous and updating if needed
func (c *cache) UpsertProduct(ctx context.Context, p models.Product) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "CacheUpsertProduct")
	defer span.Finish()

	key := fmt.Sprintf("product%v", p.ID)
	value, err := json.Marshal(p)
	if err != nil {
		return fmt.Errorf("failed to marshall product, err: <%v>", err)
	}

	return c.client.Set(&memcache.Item{
		Key:   key,
		Value: value,
	})
}
