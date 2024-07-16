package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"strconv"
	"time"
)

type CacheRepository struct {
	RedisClient *redis.Client
}

func (r *CacheRepository) Set2FA(username, code string) error {
	c := context.Background()
	if err := r.RedisClient.Set(c, "2fa:"+username, code, time.Minute*2).Err(); err != nil {
		return err
	}
	return nil
}

func (r *CacheRepository) Get2FA(username string) (string, error) {
	c := context.Background()
	return r.RedisClient.Get(c, "2fa:"+username).Result()
}

func (r *CacheRepository) SetFailedCount(username string, count int) error {
	c := context.Background()
	if err := r.RedisClient.Set(c, "fc:"+username, count, time.Minute*5).Err(); err != nil {
		return err
	}
	return nil
}

func (r *CacheRepository) GetFailedCount(username string) (int, error) {
	c := context.Background()
	result, err := r.RedisClient.Get(c, "fc:"+username).Result()
	value, err := strconv.Atoi(result)
	if err != nil {
		return 0, err
	}
	return value, nil
}

func (r *CacheRepository) SetLastFailed(username string, last time.Time) error {
	c := context.Background()
	lastStr := last.Format(time.RFC3339)
	if err := r.RedisClient.Set(c, "lf:"+username, lastStr, time.Minute*10).Err(); err != nil {
		return err
	}
	return nil
}

func (r *CacheRepository) GetLastFailed(username string) (time.Time, error) {
	c := context.Background()
	lastStr, err := r.RedisClient.Get(c, "lf:"+username).Result()
	if err != nil {
		return time.Time{}, err
	}
	lastTime, err := time.Parse(time.RFC3339, lastStr)
	if err != nil {
		return time.Time{}, err
	}
	return lastTime, nil
}
