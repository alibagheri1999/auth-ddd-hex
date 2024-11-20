package cache

import (
	"context"
	"time"
)

type CacheRepository interface {
	Set2FA(ctx context.Context, username, code string) error
	Get2FA(ctx context.Context, username string) (string, error)
	SetFailedCount(ctx context.Context, username string, count int) error
	GetFailedCount(ctx context.Context, username string) int
	SetLastFailed(ctx context.Context, username string, last time.Time) error
	GetLastFailed(ctx context.Context, username string) time.Time
	Set(ctx context.Context, key, value string, ttl uint32) error
	Get(ctx context.Context, key string) (string, error)
}
