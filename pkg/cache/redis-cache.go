package cache

import (
	"fmt"
	"time"

	redis "github.com/go-redis/redis/v8"
	"github.com/mj-hagonoy/githubprofilessvc/pkg/errors"
)

type redisCache struct {
	host     string
	password string
	db       int
	expiry   time.Duration
}

func NewRedisCache(host string, password string, db int, expiry time.Duration) *redisCache {
	return &redisCache{
		host:     host,
		password: password,
		db:       db,
		expiry:   expiry,
	}
}

func (cache *redisCache) getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cache.host,
		Password: cache.password,
		DB:       cache.db,
	})
}

func (cache *redisCache) Set(key string, value interface{}) {
	client := cache.getClient()
	err := client.Set(client.Context(), key, value, cache.expiry).Err()
	if err != nil {
		errors.Send(fmt.Errorf("redis cache error %v", err))
	}
}

func (cache *redisCache) Get(key string) *string {
	client := cache.getClient()
	cmd := client.Get(client.Context(), key)
	if cmd.Err() != nil {
		errors.Send(fmt.Errorf("redis cache error %v", cmd.Err()))
		return nil
	}
	val := cmd.Val()
	return &val
}
