package cache

import "time"

type Cache struct {
	storage map[string]CachedValue
}

type CachedValue struct {
	value     string
	timestamp int64
}

func NewCache() Cache {
	return Cache{
		storage: make(map[string]CachedValue),
	}
}

func (cache *Cache) Get(key string) (string, bool) {
	cache.ClearOld()

	var cachedValue, ok = cache.storage[key]

	if !ok {
		return "", false
	}

	return cachedValue.value, true
}

func (cache *Cache) Put(key, value string) {
	var zeroTime time.Time = time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC)
	cache.PutTill(key, value, zeroTime)
}

func (cache *Cache) Keys() []string {
	cache.ClearOld()

	var keys = make([]string, 0, len(cache.storage))

	for key := range cache.storage {
		keys = append(keys, key)
	}

	return keys
}

func (cache *Cache) ClearOld() {
	var timestampNow = time.Now().UTC().Unix()

	for key, cachedValue := range cache.storage {
		if cachedValue.timestamp > 0 && cachedValue.timestamp < timestampNow {
			delete(cache.storage, key)
		}
	}
}

func (cache *Cache) PutTill(key, value string, deadline time.Time) {
	cache.ClearOld()

	if deadline.IsZero() {
		cache.storage[key] = CachedValue{
			value:     value,
			timestamp: 0,
		}
	} else {
		var timestampNow = time.Now().UTC().Unix()

		cache.storage[key] = CachedValue{
			value:     value,
			timestamp: timestampNow,
		}
	}
}
