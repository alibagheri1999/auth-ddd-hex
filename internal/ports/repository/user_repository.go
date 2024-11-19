package repository

import (
	"DDD-HEX/internal/domain"
	"context"
)

type UserRepository interface {
	Save(ctx context.Context, user domain.UserEntity) error
	FindByID(ctx context.Context, id string) (domain.UserEntity, error)
	FindByEmail(ctx context.Context, email string) (domain.UserEntity, error)
}
