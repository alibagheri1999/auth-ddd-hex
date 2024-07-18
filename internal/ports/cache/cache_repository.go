package cache

import (
	"time"
)

type CacheRepository interface {
	Set2FA(username, code string) error
	Get2FA(username string) (string, error)
	SetFailedCount(username string, count int) error
	GetFailedCount(username string) int
	SetLastFailed(username string, last time.Time) error
	GetLastFailed(username string) time.Time
}
