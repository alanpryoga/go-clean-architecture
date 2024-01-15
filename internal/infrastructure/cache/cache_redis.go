package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisCache struct {
	db *redis.Client
}

func NewRedisCache(db *redis.Client) Cache {
	return &RedisCache{
		db: db,
	}
}

func (c *RedisCache) Get(ctx context.Context, key string) ([]byte, error) {
	val, err := c.db.Get(ctx, key).Bytes()
	if err == redis.Nil {
		return nil, ErrKeyNotFound
	} else if err != nil {
		return nil, err
	}
	return val, nil
}

func (c *RedisCache) Set(ctx context.Context, key string, value []byte) error {
	return c.db.Set(ctx, key, value, 0).Err()
}

func (c *RedisCache) SetWithTTL(ctx context.Context, key string, value []byte, ttl time.Duration) error {
	return c.db.SetEx(ctx, key, value, ttl).Err()
}

func (c *RedisCache) Delete(ctx context.Context, key string) error {
	return c.db.Del(ctx, key).Err()
}

func (c *RedisCache) Ping(ctx context.Context) error {
	_, err := c.db.Ping(ctx).Result()
	return err
}
