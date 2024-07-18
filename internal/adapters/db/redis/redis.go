package redis

import (
	"DDD-HEX/config"
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

var ctx = context.Background()

type ClientWrapper struct {
	Client *redis.Client
}

func NewRedisClientWrapper(config config.RedisConfig) *ClientWrapper {
	client, err := NewRedisClient(config)
	if err != nil {
		logrus.Error("Failed to connect to Redis with err ", err.Error())
		os.Exit(1)
	}
	return &ClientWrapper{
		Client: client,
	}
}

func NewRedisClient(config config.RedisConfig) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     config.Host,
		Password: config.Password,
		DB:       config.DB,
	})
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := client.Ping(ctx).Result()
	if err != nil {
		logrus.Warn("Failed to connect to Redis:", err)
		return nil, err
	}
	logrus.Info("Successfully connected to Redis")
	return client, nil
}

func (w *ClientWrapper) EnsureConnected(maxRetries int) {
	ctx := context.Background()
	_, err := w.Client.Ping(ctx).Result()
	if err != nil {
		logrus.Warn("Lost connection to Redis, attempting to reconnect...")
		if !RetryRedisConnection(w.Client, maxRetries) {
			logrus.Warn("Could not reconnect to Redis after max retries, exiting...")
			return
		}
	}
}

func RetryRedisConnection(client *redis.Client, maxRetries int) bool {
	var attempt int
	var err error

	for attempt = 0; attempt < maxRetries; attempt++ {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		_, err = client.Ping(ctx).Result()
		if err == nil {
			logrus.Info("Successfully reconnected to Redis")
			return true
		}

		time.Sleep(time.Duration(attempt+1) * time.Second)
	}

	logrus.Warn("Failed to reconnect to Redis after", maxRetries, "attempts:", err)
	return false
}
