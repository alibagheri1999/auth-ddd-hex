package redis

import (
	"context"
	"strconv"
	"time"
)

type CacheRepository struct {
	RedisClient *ClientWrapper
}

func (r *CacheRepository) Set2FA(username, code string) error {
	c := context.Background()
	if err := r.RedisClient.Client.Set(c, "2fa:"+username, code, time.Minute*2).Err(); err != nil {
		return err
	}
	return nil
}

func (r *CacheRepository) Get2FA(username string) (string, error) {
	c := context.Background()
	return r.RedisClient.Client.Get(c, "2fa:"+username).Result()
}

func (r *CacheRepository) SetFailedCount(username string, count int) error {
	c := context.Background()
	if err := r.RedisClient.Client.Set(c, "fc:"+username, count, time.Minute*5).Err(); err != nil {
		return err
	}
	return nil
}

func (r *CacheRepository) GetFailedCount(username string) int {
	c := context.Background()
	result, err := r.RedisClient.Client.Get(c, "fc:"+username).Result()
	if err != nil {
		return 0
	}
	value, err := strconv.Atoi(result)
	if err != nil {
		return 0
	}
	return value
}

func (r *CacheRepository) SetLastFailed(username string, last time.Time) error {
	c := context.Background()
	lastStr := last.Format(time.RFC3339)
	if err := r.RedisClient.Client.Set(c, "lf:"+username, lastStr, time.Minute*10).Err(); err != nil {
		return err
	}
	return nil
}

func (r *CacheRepository) GetLastFailed(username string) time.Time {
	c := context.Background()
	lastStr, err := r.RedisClient.Client.Get(c, "lf:"+username).Result()
	if err != nil {
		return time.Time{}
	}
	lastTime, err := time.Parse(time.RFC3339, lastStr)
	if err != nil {
		return time.Time{}
	}
	return lastTime
}
