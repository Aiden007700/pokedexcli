package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	mu    sync.Mutex
	Cache map[string]cacheEntry
	debug bool
}

func (c *Cache) Add(key string, val []byte) {
	c.Cache[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	if val, ok := c.Cache[key]; ok {
		return val.val, ok
	} else {
		return nil, ok
	}
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)

	go func() {
		for {
			select {
			case t := <-ticker.C:
				if c.debug {
					fmt.Println("Tick at", t)
				}
				c.mu.Lock()
				for k, v := range c.Cache {
					if time.Since(v.createdAt) > interval {
						if c.debug {
							fmt.Println(c.Cache[k])
							fmt.Println("Deleting from cache", k)
						}
						delete(c.Cache, k)
						if c.debug {
							fmt.Println(c.Cache[k])
						}
					}
				}
				c.mu.Unlock()
			}
		}
	}()
}

func NewCache(interval time.Duration, debug bool) *Cache {
	cache := &Cache{
		mu:    sync.Mutex{},
		Cache: make(map[string]cacheEntry),
		debug: debug,
	}

	cache.reapLoop(interval)

	return cache
}
