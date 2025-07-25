package db

import "github.com/Cenzios/pf-backend/pkg/db/dbiface"

// CacheStore is the global cache instance
var CacheStore dbiface.Cache

// Cache interface is now in dbiface
// type Cache = dbiface.Cache

func FindOneCache(key string) (string, error) {
	return CacheStore.Get(key)
}

func SetCache(key string, value string, ttlSeconds int) error {
	return CacheStore.Set(key, value, ttlSeconds)
}

func DeleteCache(key string) error {
	return CacheStore.Delete(key)
}

func CacheExists(key string) (bool, error) {
	return CacheStore.Exists(key)
}
