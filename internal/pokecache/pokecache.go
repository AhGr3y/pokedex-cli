package pokecache

import (
	"errors"
	"sync"
	"time"
)

type Cache struct {
	entries  map[string]cacheEntry
	mu       *sync.RWMutex
	interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		entries:  map[string]cacheEntry{},
		mu:       &sync.RWMutex{},
		interval: interval,
	}

	go cache.reapLoop()

	return cache
}

func (c *Cache) Add(url string, data []byte) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if url == "" {
		return errors.New("unable to add entry: missing url")
	}

	if data == nil {
		return errors.New("unable to add entry: missing data")
	}

	c.entries[url] = cacheEntry{
		createdAt: time.Now(),
		val:       data,
	}

	return nil
}

func (c *Cache) Get(url string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if url == "" {
		return nil, false
	}

	data, ok := c.entries[url]
	if !ok {
		return nil, false
	}

	return data.val, true
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	for range ticker.C {
		for key, cacheEntry := range c.entries {
			elapsed := time.Since(cacheEntry.createdAt)
			if elapsed >= c.interval {
				c.mu.Lock()
				delete(c.entries, key)
				c.mu.Unlock()
			}
		}
	}

}
