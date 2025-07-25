package memcache

import (
	"os"
	"strings"

	"github.com/Cenzios/pf-backend/pkg/db/dbiface"
	"github.com/bradfitz/gomemcache/memcache"
)

// memcacheImpl implements dbiface.Cache
// It wraps a memcache.Client

type memcacheImpl struct {
	client *memcache.Client
}

// New returns a new memcacheImpl as dbiface.Cache
func New() dbiface.Cache {
	client := memcache.New(strings.Split(os.Getenv("MEMCACHE_ADDR"), ",")...)
	return &memcacheImpl{client}
}

// Get retrieves a value by key
func (m *memcacheImpl) Get(key string) (string, error) {
	item, err := m.client.Get(key)
	if err != nil {
		return "", err
	}
	return string(item.Value), nil
}

// Set sets a value with a TTL in seconds
func (m *memcacheImpl) Set(key, value string, ttlSeconds int) error {
	return m.client.Set(&memcache.Item{
		Key:        key,
		Value:      []byte(value),
		Expiration: int32(ttlSeconds),
	})
}

// Delete removes a key
func (m *memcacheImpl) Delete(key string) error {
	return m.client.Delete(key)
}

// Exists checks if a key exists
func (m *memcacheImpl) Exists(key string) (bool, error) {
	_, err := m.client.Get(key)
	return err == nil, err
}
