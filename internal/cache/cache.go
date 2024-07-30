package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisCache struct {
	client *redis.Client
}

func NewRedisCache(addr string) *RedisCache {
	rdb := redis.NewClient(&redis.Options{
		Addr: addr,
	})
	return &RedisCache{client: rdb}
}

func (c *RedisCache) Get(ctx context.Context, key string, dest interface{}) error {
	res, err := c.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil
	} else if err != nil {
		return err
	}

	return json.Unmarshal([]byte(res), dest)
}

func (c *RedisCache) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return c.client.Set(ctx, key, data, ttl).Err()
}
