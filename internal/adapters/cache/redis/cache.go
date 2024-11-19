package redis

import (
	"DDD-HEX/internal/ports/clients"
	"context"
	"strconv"
	"time"
)

type CacheRepository struct {
	Cache clients.Cache
}

func (r *CacheRepository) EnsureConnected(maxRetries int) error {
	return r.Cache.EnsureConnected(maxRetries)
}

func (r *CacheRepository) Set2FA(ctx context.Context, username, code string) error {
	key := "2fa:" + username
	return r.Cache.Set(ctx, key, code, 120) // 2 minutes TTL
}

func (r *CacheRepository) Get2FA(ctx context.Context, username string) (string, error) {
	key := "2fa:" + username
	val, err := r.Cache.Get(ctx, key)
	if err != nil {
		return "", err
	}
	return val, nil
}

func (r *CacheRepository) SetFailedCount(ctx context.Context, username string, count int) error {
	key := "fc:" + username
	return r.Cache.Set(ctx, key, count, 300) // 5 minutes TTL
}

func (r *CacheRepository) GetFailedCount(ctx context.Context, username string) int {
	key := "fc:" + username
	val, err := r.Cache.Get(ctx, key)
	if err != nil {
		return 0 // Default if key is missing
	}
	count, err := strconv.Atoi(val)
	if err != nil {
		return 0
	}
	return count
}

func (r *CacheRepository) SetLastFailed(ctx context.Context, username string, last time.Time) error {
	key := "lf:" + username
	lastStr := last.Format(time.RFC3339)
	return r.Cache.Set(ctx, key, lastStr, 600) // 10 minutes TTL
}

func (r *CacheRepository) GetLastFailed(ctx context.Context, username string) time.Time {
	key := "lf:" + username
	val, err := r.Cache.Get(ctx, key)
	if err != nil {
		return time.Time{} // Default if key is missing
	}
	lastTime, err := time.Parse(time.RFC3339, val)
	if err != nil {
		return time.Time{}
	}
	return lastTime
}
