package config

import "github.com/go-redis/redis/v8"

type (
	CacheStruct struct {
		Options Options
		client  *redis.Client
	}

	Options struct {
		Address     string   `yaml:"address"`
		Password    string   `yaml:"password"`
		DB          int      `yaml:"db"`
		UseSentinel bool     `yaml:"use_sentinel"`
		Sentinel    Sentinel `yaml:"sentinel"`
	}

	Sentinel struct {
		Name  string   `yaml:"name"`
		Addrs []string `yaml:"addresses"`
	}

	KeyNotFoundError struct{}

	ConnectionError struct {
		Message string
	}
)
