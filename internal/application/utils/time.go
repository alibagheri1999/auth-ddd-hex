package utils

import (
	"time"
)

func CalculateTimeDifference(storedTime time.Time) time.Duration {
	now := time.Now()
	return now.Sub(storedTime)
}

func IsWithinTenMinutes(duration time.Duration) bool {
	return duration.Minutes() >= 10
}
