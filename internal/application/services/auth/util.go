package auth

import (
	"DDD-HEX/internal/application/utils"
	"DDD-HEX/internal/ports/repository"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

const (
	charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
)

func GenerateRandomCode(length int) string {
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	code := make([]byte, length)
	for i := range code {
		code[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(code)
}

func HandleFailLogin(email string, cacheRepo repository.CacheRepository) (int, error) {
	failedCount, err := cacheRepo.GetFailedCount(email)
	if err != nil {
		return 0, err
	}
	lastFailed, err := cacheRepo.GetLastFailed(email)
	if err != nil {
		return 0, err
	}
	duration := utils.CalculateTimeDifference(lastFailed)
	text := fmt.Sprintf("please try again after %f mins, you rached 3 fail tries", duration.Minutes())
	if failedCount == 3 {
		if sErr := cacheRepo.SetFailedCount(email, 0); sErr != nil {
			return 0, sErr
		}
		if sErr := cacheRepo.SetLastFailed(email, time.Now()); sErr != nil {
			return 0, sErr
		}
		text := fmt.Sprintf("please try again after %d mins, you rached 3 fail tries", 10)
		return 0, errors.New(text)
	}
	if !utils.IsWithinTenMinutes(duration) {
		return 0, errors.New(text)
	}
	return failedCount, nil
}
