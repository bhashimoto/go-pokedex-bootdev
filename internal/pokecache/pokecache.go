package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache	map[string]cacheEntry
	mux 	*sync.Mutex
}

type cacheEntry struct {
	createdAt	time.Time
	val 		[]byte
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		cache: make(map[string]cacheEntry),
		mux: &sync.Mutex{},
	}

	go c.reapLoop(interval)

	return c
}

func (c Cache) Add(key string, val []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()

	c.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val: val,
	}
}

func (c Cache) Get(key string) (val []byte, ok bool) {
	c.mux.Lock()
	defer c.mux.Unlock()
	entry, ok := c.cache[key]
	if ok {
		return entry.val, true
	}
	return []byte{}, false 
}

func (c Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for {
		c.mux.Lock()
		for k, v := range c.cache {
			if v.createdAt.Before(time.Now().Add(-interval)) {
				delete(c.cache, k)
			}
		}
		c.mux.Unlock()
	}
}
