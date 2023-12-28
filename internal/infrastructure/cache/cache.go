package cache

import (
	"context"
	"time"
)

//go:generate mockgen -source=cache.go -destination=cache_mock.go -package=cache

// Cache is an interface that defines methods for a generic caching mechanism.
type Cache interface {
	// Get retrieves the value associated with the given key from the cache.
	Get(ctx context.Context, key string) ([]byte, error)

	// Set stores the specified value in the cache with the given key.
	Set(ctx context.Context, key string, value []byte) error

	// SetWithTTL stores the specified value in the cache with the given key and a time-to-live (TTL) duration.
	// The value will be automatically evicted from the cache after the specified TTL duration.
	SetWithTTL(ctx context.Context, key string, value []byte, ttl time.Duration) error

	// Delete removes the entry associated with the given key from the cache.
	Delete(ctx context.Context, key string) error

	// Ping checks the availability of the cache.
	Ping(ctx context.Context) error
}
