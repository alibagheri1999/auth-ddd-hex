package repository

import (
	"DDD-HEX/internal/domain"
	"context"
)

type AuthRepository interface {
	Save(ctx context.Context, auth domain.AuthEntity) error
	FindByAccessToken(ctx context.Context, token string) (domain.AuthEntity, error)
	FindByRefreshToken(ctx context.Context, token string) (domain.UserAuthEntity, error)
}
