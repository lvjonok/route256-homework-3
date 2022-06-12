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

func (c *cache) GetReviews(ctx context.Context, id types.ID) ([]models.Review, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "CacheGetProduct")
	defer span.Finish()

	item, err := c.client.Get(fmt.Sprintf("reviews%v", id))
	if errors.Is(err, memcache.ErrCacheMiss) {
		return nil, cacheconnector.ErrCacheMiss
	} else if err != nil {
		return nil, fmt.Errorf("failed to get reviews, err: <%v>", err)
	}

	var reviews []models.Review
	if err := json.Unmarshal(item.Value, &reviews); err != nil {
		return nil, fmt.Errorf("failed to unmarshall product, err: <%v>", err)
	}

	return reviews, nil
}

// AppendReviews breaks our atomicity and we just remove the entry with revies for this product from cache, to fetch it again, but later
func (c *cache) AppendReview(ctx context.Context, r models.Review) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "CacheAppendReview")
	defer span.Finish()

	key := fmt.Sprintf("reviews%v", r.ProductID)
	if err := c.client.Delete(key); err != memcache.ErrCacheMiss {
		return fmt.Errorf("failed to delete key, err: <%v>", err)
	}
	return nil
}

func (c *cache) UpsertReviews(ctx context.Context, productID types.ID, reviews []models.Review) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "CacheUpsertReviews")
	defer span.Finish()

	key := fmt.Sprintf("reviews%v", productID)
	value, err := json.Marshal(reviews)
	if err != nil {
		return fmt.Errorf("failed to marshall product, err: <%v>", err)
	}

	return c.client.Set(&memcache.Item{
		Key:   key,
		Value: value,
	})
}
