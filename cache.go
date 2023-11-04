package go_ernie

import (
	"sync"
	"time"
)

type Cache struct {
	data map[string]cacheItem
	lock sync.RWMutex
}

type cacheItem struct {
	value      interface{}
	expiration time.Time
}

func NewCache() *Cache {
	return &Cache{
		data: make(map[string]cacheItem),
	}
}

func (c *Cache) Get(key string) (interface{}, bool) {
	c.lock.RLock()
	defer c.lock.RUnlock()
	item, exists := c.data[key]
	if !exists {
		return nil, false
	}
	// 检查过期时间
	if time.Now().After(item.expiration) {
		// 过期了，从缓存中删除该项
		delete(c.data, key)
		return nil, false
	}
	return item.value, true
}

func (c *Cache) Set(key string, value interface{}, expiration time.Duration) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.data[key] = cacheItem{
		value:      value,
		expiration: time.Now().Add(expiration),
	}
}

func (c *Cache) Del(key string) {
	c.lock.Lock()
	defer c.lock.Unlock()
	_, exists := c.data[key]
	if exists {
		delete(c.data, key)
	}

}
