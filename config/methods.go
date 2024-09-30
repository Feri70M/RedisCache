package config

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

func (c *CacheStruct) Set(key, val string, ttl time.Duration) error {
	return c.client.Set(context.Background(), key, val, ttl).Err()
}

func (c *CacheStruct) Get(key string) (string, error) {
	val, err := c.client.Get(context.Background(), key).Result()
	if err == redis.Nil {
		return "", KeyNotFoundError{}
	} else if err != nil {
		return "", err
	}
	return val, nil
}

func (c *CacheStruct) Delete(key string) error {
	return c.client.Del(context.Background(), key).Err()
}

func (c *CacheStruct) Exists(key string) (bool, error) {
	result, err := c.client.Exists(context.Background(), key).Result()
	if err != nil {
		return false, err
	}
	return result > 0, nil
}

func (c *CacheStruct) Close() error {
	return c.client.Close()
}

func (e KeyNotFoundError) Error() string {
	return "key not found"
}
