package auth

import (
	"DDD-HEX/internal/domain"
	"DDD-HEX/internal/domain/DTO"
	"context"
	"github.com/stretchr/testify/mock"
	"time"
)

// MockAuthRepository
type MockAuthRepository struct {
	mock.Mock
}

func (m *MockAuthRepository) Save(ctx context.Context, auth domain.AuthEntity) error {
	args := m.Called(ctx, auth)
	return args.Error(0)
}

func (m *MockAuthRepository) FindByRefreshToken(ctx context.Context, token string) (domain.UserAuthEntity, error) {
	args := m.Called(ctx, token)
	return args.Get(0).(domain.UserAuthEntity), args.Error(1)
}

func (m *MockAuthRepository) FindByAccessToken(ctx context.Context, token string) (domain.AuthEntity, error) {
	args := m.Called(ctx, token)
	return args.Get(0).(domain.AuthEntity), args.Error(1)
}

// MockUserService
type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) FindUserByEmail(ctx context.Context, email string) (*domain.UserEntity, error) {
	args := m.Called(ctx, email)
	return args.Get(0).(*domain.UserEntity), args.Error(1)
}

func (m *MockUserService) CreateUser(c context.Context, dto DTO.CreateUserRequest) error {
	args := m.Called(c, dto)
	return args.Error(0)
}

// MockCacheRepository
type MockCacheRepository struct {
	mock.Mock
}

func (m *MockCacheRepository) SetFailedCount(ctx context.Context, email string, count int) error {
	args := m.Called(ctx, email, count)
	return args.Error(0)
}

func (m *MockCacheRepository) Set2FA(ctx context.Context, username, code string) error {
	args := m.Called(ctx, username, code)
	return args.Error(0)
}

func (m *MockCacheRepository) Get2FA(ctx context.Context, username string) (string, error) {
	args := m.Called(ctx, username)
	return args.String(0), args.Error(1)
}

func (m *MockCacheRepository) Set(ctx context.Context, key, value string, ttl uint32) error {
	args := m.Called(ctx, key, value, ttl)
	return args.Error(0)
}

func (m *MockCacheRepository) Get(ctx context.Context, key string) (string, error) {
	args := m.Called(ctx, key)
	return args.String(0), args.Error(1)
}

// Add this method to mock GetFailedCount
func (m *MockCacheRepository) GetFailedCount(ctx context.Context, username string) int {
	args := m.Called(ctx, username)
	return args.Int(0)
}

// Mock other cache methods
func (m *MockCacheRepository) SetLastFailed(ctx context.Context, username string, last time.Time) error {
	args := m.Called(ctx, username, last)
	return args.Error(0)
}

func (m *MockCacheRepository) GetLastFailed(ctx context.Context, username string) time.Time {
	args := m.Called(ctx, username)
	return args.Get(0).(time.Time)
}
