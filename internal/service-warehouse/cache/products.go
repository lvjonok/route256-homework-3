package cache

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/opentracing/opentracing-go"
	types "gitlab.ozon.dev/lvjonok/homework-3/core/models"
)

// GetProducts tries to get all products by ids from cache
func (c *cache) GetProducts(ctx context.Context, ids []types.ID) ([]types.ProductUnit, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "CacheGetProducts")
	defer span.Finish()

	keys := []string{}
	for _, i := range ids {
		keys = append(keys, fmt.Sprintf("product%v", i))
	}

	values, err := c.client.GetMulti(keys)
	if err != nil {
		return nil, fmt.Errorf("failed to get products from cache, err: <%v>", err)
	}

	units := []types.ProductUnit{}
	for _, v := range values {
		var unit types.ProductUnit
		if err := json.Unmarshal(v.Value, &unit); err != nil {
			return nil, fmt.Errorf("failed to unmarshall product, err: <%v>", err)
		}

		units = append(units, unit)
	}

	return units, nil
}

// UpsertProducts updates values for products in cache
func (c *cache) UpsertProducts(ctx context.Context, units []types.ProductUnit) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "CacheUpsertProducts")
	defer span.Finish()

	for _, u := range units {
		key := fmt.Sprintf("product%v", u.ProductID)
		data, err := json.Marshal(u)
		if err != nil {
			return fmt.Errorf("failed to marshal product, err: <%v>", err)
		}

		if err := c.client.Set(&memcache.Item{
			Key:   key,
			Value: data,
		}); err != nil {
			return fmt.Errorf("failed to set product in cache, err: <%v>", err)
		}
	}

	return nil
}

// DeleteProducts removes ids of products from cache, so we would refresh them
func (c *cache) DeleteProducts(ctx context.Context, ids []types.ID) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "CacheDeleteProducts")
	defer span.Finish()

	for _, i := range ids {
		key := fmt.Sprintf("product%v", i)
		if err := c.client.Delete(key); err != nil && err != memcache.ErrCacheMiss {
			return fmt.Errorf("failed to delete key, err: <%v>", err)
		}
	}

	return nil
}
