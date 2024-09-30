package config

import "time"

type CacheInterface interface {
	Set(key, val string, ttl time.Duration) error
	Get(key string) (string, error)
	Delete(key string) error
	Exists(key string) (bool, error)
	Close() error
}
