package auth

import (
	"DDD-HEX/internal/application/utils"
	"DDD-HEX/internal/ports/cache"
	"context"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"math"
	"math/rand"
	"time"
)

const (
	charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
)

var (
	Secret = []byte(utils.ConfigSetup().App.JwtSecret)
)

type Claims struct {
	Email  string `json:"eamil"`
	Role   string `json:"role"`
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

func GenerateRandomCode(length int) string {
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	code := make([]byte, length)
	for i := range code {
		code[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(code)
}

func HandleFailLogin(ctx context.Context, email string, cacheRepo cache.CacheRepository) (int, error) {
	failedCount := cacheRepo.GetFailedCount(ctx, email)
	lastFailed := cacheRepo.GetLastFailed(ctx, email)
	duration := utils.CalculateTimeDifference(lastFailed)
	text := fmt.Sprintf("please try again after %v mins, you rached 3 fail tries", 10-math.Floor(duration.Minutes()))
	if failedCount == 3 {
		if sErr := cacheRepo.SetFailedCount(ctx, email, 0); sErr != nil {
			return 0, sErr
		}
		if sErr := cacheRepo.SetLastFailed(ctx, email, time.Now()); sErr != nil {
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
