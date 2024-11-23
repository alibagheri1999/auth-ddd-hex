package auth_test

import (
	"DDD-HEX/config"
	"DDD-HEX/internal/application/services/auth"
	"DDD-HEX/internal/application/utils"
	"DDD-HEX/internal/domain"
	"context"
	"database/sql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

// TestAuthenticate_Success with mocked CheckHash
// Test for Authenticate function with mocked CheckHash
func TestAuthenticate_Success(t *testing.T) {
	// Setup mocks
	authRepo := new(auth.MockAuthRepository)
	userService := new(auth.MockUserService)
	cacheRepo := new(auth.MockCacheRepository)

	// Create a mock user
	user := &domain.UserEntity{
		ID:       1,
		Email:    "test2@example.com",
		Password: sql.NullString{String: "hashedpassword", Valid: true},
		Status:   "active",
	}

	// Mock method to find user by email
	userService.On("FindUserByEmail", mock.Anything, "test2@example.com").Return(user, nil)

	// Mock the GetFailedCount to return 0 (indicating no failed login attempts)
	cacheRepo.On("GetFailedCount", mock.Anything, "test2@example.com").Return(0)

	// Mock GetLastFailed to return a fixed time (or any value based on your logic)
	cacheRepo.On("GetLastFailed", mock.Anything, "test2@example.com").Return(time.Now()) // or time.Time{}

	// Override CheckHash function for this test
	originalCheckHash := utils.CheckHash
	defer func() { utils.CheckHash = originalCheckHash }() // Restore original after test

	// Mock password validation to always return true
	utils.CheckHash = func(password, hashedPassword string) bool {
		return password == "correctpassword"
	}

	// Mock generating tokens
	authService := auth.NewAuthService(authRepo, userService, cacheRepo, config.AppConfig{}, utils.DefaultCheckHash)

	// Mock Save method to just return nil
	authRepo.On("Save", mock.Anything, mock.Anything).Return(nil)

	// Call the Authenticate method
	accessToken, refreshToken, err := authService.Authenticate(context.Background(), "test2@example.com", "correctpassword")

	// Assertions
	assert.NoError(t, err)
	assert.NotEmpty(t, accessToken)
	assert.NotEmpty(t, refreshToken)
}
