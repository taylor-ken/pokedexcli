package pokecache

import (
	"sync"
	"time"
)

// Cache -
type Cache struct {
	entries map[string]cacheEntry
	mux     sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(timeout time.Duration) *Cache {
	cache := &Cache{
		entries: make(map[string]cacheEntry),
		mux:     sync.Mutex{},
	}
	go cache.reapLoop(timeout)
	return cache
}

func (cache *Cache) Add(key string, val []byte) {
	cache.mux.Lock()
	defer cache.mux.Unlock()
	now := time.Now()
	cache.entries[key] = cacheEntry{
		createdAt: now,
		val:       val,
	}
}

func (cache *Cache) Get(key string) ([]byte, bool) {
	cache.mux.Lock()
	defer cache.mux.Unlock()
	entry, ok := cache.entries[key]
	if ok {
		return entry.val, true
	}
	return []byte{}, false
}

func (cache *Cache) reapLoop(timeout time.Duration) {
	ticker := time.NewTicker(timeout)
	defer ticker.Stop()
	for range ticker.C {
		cache.mux.Lock()
		for key, val := range cache.entries {
			if time.Since(val.createdAt) > timeout {
				delete(cache.entries, key)
			}
		}
		cache.mux.Unlock()
	}
}
