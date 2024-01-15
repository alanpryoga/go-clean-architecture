package cache

import (
	"context"
	"errors"
	"time"
)

//go:generate mockgen -source=cache.go -destination=cache_mock.go -package=cache

var (
	ErrKeyNotFound = errors.New("key is not exists")
)

type Cache interface {
	Get(ctx context.Context, key string) ([]byte, error)
	Set(ctx context.Context, key string, value []byte) error
	SetWithTTL(ctx context.Context, key string, value []byte, ttl time.Duration) error
	Delete(ctx context.Context, key string) error
	Ping(ctx context.Context) error
}
