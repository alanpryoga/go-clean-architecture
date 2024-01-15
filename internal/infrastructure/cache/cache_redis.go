package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type redisCache struct {
	client *redis.Client
}

func NewRedisCache(db *redis.Client) Cache {
	return &redisCache{
		client: db,
	}
}

func (c *redisCache) Get(ctx context.Context, key string) ([]byte, error) {
	val, err := c.client.Get(ctx, key).Bytes()
	if err == redis.Nil {
		return nil, ErrKeyNotFound
	} else if err != nil {
		return nil, err
	}
	return val, nil
}

func (c *redisCache) Set(ctx context.Context, key string, value []byte) error {
	return c.client.Set(ctx, key, value, 0).Err()
}

func (c *redisCache) SetWithTTL(ctx context.Context, key string, value []byte, ttl time.Duration) error {
	return c.client.SetEx(ctx, key, value, ttl).Err()
}

func (c *redisCache) Delete(ctx context.Context, key string) error {
	return c.client.Del(ctx, key).Err()
}

func (c *redisCache) Ping(ctx context.Context) error {
	_, err := c.client.Ping(ctx).Result()
	return err
}
