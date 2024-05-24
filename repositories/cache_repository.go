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

func (c *CacheRepository) Set(referenceId string, otp string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[referenceId] = otp
}

func (c *CacheRepository) Get(referenceId string) (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	otp, exists := c.data[referenceId]
	return otp, exists
}
