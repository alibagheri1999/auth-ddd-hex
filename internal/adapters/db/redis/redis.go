package redis

import (
	"DDD-HEX/config"
	"context"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func NewRedisClient(config config.RedisConfig) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     config.Host,
		Password: config.Password,
		DB:       config.DB,
	})
}
