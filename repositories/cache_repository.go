package repositories

import "sync"

type CacheRepository interface {
	Set(name, val string)
	Get(name string) (string, bool)
}

type cacheRepository struct {
	data map[string]string
	mu   sync.RWMutex
}

func NewCacheRepository() cacheRepository {
	return cacheRepository{
		data: make(map[string]string),
	}
}

func (c *cacheRepository) Set(name, val string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[name] = val
}

func (c *cacheRepository) Get(name string) (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	res, exists := c.data[name]
	return res, exists
}
