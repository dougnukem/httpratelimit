package httpratelimit

import (
	"github.com/patrickmn/go-cache"
	"time"
)

type Cache struct {
	cache *cache.Cache
}

func NewDefaultCache() *Cache {
	return NewCache(5*time.Minute, 30*time.Second)
}

func NewCache(defaultExpiration, cleanupInterval time.Duration) *Cache {
	return &Cache{cache: cache.New(defaultExpiration, cleanupInterval)}
}

func (c *Cache) Store(key string, value []byte, timeout time.Duration) error {
	c.cache.Set(key, value, timeout)
	return nil
}

func (c *Cache) Get(key string) ([]byte, error) {
	v, found := c.cache.Get(key)
	if !found {
		return nil, nil
	} else {
		return v.([]byte), nil
	}
}
