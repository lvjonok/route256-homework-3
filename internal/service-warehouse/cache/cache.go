package cache

import (
	"github.com/bradfitz/gomemcache/memcache"
)

type cache struct {
	client *memcache.Client
}

func New(urls []string) (*cache, error) {
	mc := memcache.New(urls...)

	return &cache{
		client: mc,
	}, mc.Ping()
}
