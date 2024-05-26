package cache

import "sync"

type Cache interface {
	Get(key string) (interface{}, bool)
	Set(key string, value interface{})
	Delete(key string)
}

type MemoryCache struct {
	data map[string]interface{}
	mu   sync.RWMutex
}

func NewMemoryCache() *MemoryCache {
	return &MemoryCache{
		data: make(map[string]interface{}),
	}
}

func (c *MemoryCache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	value, ok := c.data[key]
	return value, ok
}

func (c *MemoryCache) Set(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = value
}

func (c *MemoryCache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.data, key)
}
