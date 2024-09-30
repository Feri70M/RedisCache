package config

import (
	"errors"

	"github.com/go-redis/redis/v8"
)

func NewCache(opts Options) (CacheInterface, error) {

	if opts.UseSentinel {
		if opts.Sentinel.Name == "" || len(opts.Sentinel.Addrs) == 0 {
			return nil, errors.New("sentinel master name and addresses must be provided when using sentinel")
		}
		sentinelClient := redis.NewFailoverClient(&redis.FailoverOptions{
			MasterName:    opts.Sentinel.Name,
			SentinelAddrs: opts.Sentinel.Addrs,
			DB:            opts.DB,
			Password:      opts.Password,
		})

		return &CacheStruct{Options: opts, client: sentinelClient}, nil
	}

	standaloneClient := redis.NewClient(&redis.Options{
		Addr:     opts.Address,
		Password: opts.Password,
		DB:       opts.DB,
	})

	return &CacheStruct{Options: opts, client: standaloneClient}, nil
}
