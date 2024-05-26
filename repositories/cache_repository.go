package repositories

import "sync"

type CacheRepository struct {
	data map[string]string
	mu   sync.RWMutex
}

func NewCacheRepository() *CacheRepository {
	return &CacheRepository{
		data: make(map[string]string),
	}
}

func (c *CacheRepository) Set(name string, val string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[name] = val
}

func (c *CacheRepository) Get(name string) (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	otp, exists := c.data[name]
	return otp, exists
}