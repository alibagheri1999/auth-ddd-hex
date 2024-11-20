package redis

import (
	"DDD-HEX/config"
	"DDD-HEX/internal/ports/clients"
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"time"
)

type RedisCache struct {
	Client *redis.Client
	Config config.CacheConfig
}

func NewRedisCache(ctx context.Context, config config.CacheConfig) clients.Cache {
	client := redis.NewClient(&redis.Options{
		Addr:     config.Host,
		Password: config.Password,
		DB:       config.DB,
	})

	if _, err := client.Ping(ctx).Result(); err != nil {
		logrus.Fatalf("Failed to connect to Redis: %v", err)
	}
	return &RedisCache{
		Client: client,
		Config: config,
	}
}

func (r *RedisCache) EnsureConnected(maxRetries int) error {
	var attempt int
	var err error

	for attempt = 0; attempt < maxRetries; attempt++ {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		_, err = r.Client.Ping(ctx).Result()
		if err == nil {
			logrus.Info("Successfully connected to Redis")
			return nil
		}

		time.Sleep(time.Duration(attempt+1) * time.Second)
	}

	logrus.Warn("Failed to reconnect to Redis after", maxRetries, "attempts:", err)
	return err
}

func (r *RedisCache) Connect() error {
	_, err := r.Client.Ping(context.Background()).Result()
	return err
}

func (r *RedisCache) Close() error {
	return r.Client.Close()
}

func (r *RedisCache) Ping(ctx context.Context) error {
	_, err := r.Client.Ping(ctx).Result()
	return err
}

func (r *RedisCache) Set(ctx context.Context, key string, value interface{}, ttlSeconds uint32) error {
	return r.Client.Set(ctx, key, value, time.Duration(ttlSeconds)*time.Second).Err()
}

func (r *RedisCache) Get(ctx context.Context, key string) (string, error) {
	val, err := r.Client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", nil
	}
	return val, err
}

func (r *RedisCache) Delete(ctx context.Context, key string) error {
	return r.Client.Del(ctx, key).Err()
}

func (r *RedisCache) Flush(ctx context.Context) error {
	return r.Client.FlushAll(ctx).Err()
}
