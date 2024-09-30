package config

import "github.com/go-redis/redis/v8"

type (
	CacheStruct struct {
		Options Options
		client  *redis.Client
	}

	Options struct {
		Address     string
		Password    string
		DB          int
		UseSentinel bool
		Sentinel    Sentinel
	}

	Sentinel struct {
		Name  string
		Addrs []string
	}

	KeyNotFoundError struct{}
)
