package cache

import (
	"github.com/go-redis/redis/v8"
	"context"
)

type RedisCache struct {
	client *redis.Client
}

func NewRedisCache() (*RedisCache, error) {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	return &RedisCache{client: client}, nil
}

func (r *RedisCache) Get(key string) (string, error) {
	return r.client.Get(context.Background(), key).Result()
}

func (r *RedisCache) Set(key string, value string) error {
	return r.client.Set(context.Background(), key, value, 0).Err()
}