package main

import (
	"sync"
	"time"
	
)

type Cache struct {
	lock sync.RWMutex
	entries map[string]cacheEntry
}

type cacheEntry struct {
	createdAt time.Time
	val parsedData
}

func NewCache(interval time.Duration) Cache {
	var cache Cache
	go func() {
		cache.reapLoop(interval)
	} ()
	return cache
}

func (cache *Cache) Add(key string, val parsedData) {
	cache.lock.Lock()
	cache.entries[key] = cacheEntry { 
		createdAt: time.Time{}, 
		val : val,
	}
	cache.lock.Unlock()
}

func (cache *Cache) Get(key string) (parsedData, bool) {
	cache.lock.RLock()
	defer cache.lock.RUnlock()
	data, ok := cache.entries[key]
	return data.val, ok
}

func (cache *Cache) reapLoop(interval time.Duration) {
	ch := time.NewTicker(interval).C
	for true {
		time := <-ch
		cache.lock.Lock()
		for k,v := range cache.entries {
			if time.Sub(v.createdAt) > interval {
				delete(cache.entries, k)
			}
		}
		cache.lock.Unlock()
	}
}