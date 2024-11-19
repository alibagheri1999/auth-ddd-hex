package user

import (
	"DDD-HEX/internal/domain"
	"DDD-HEX/internal/domain/DTO"
	"context"
)

type UserService interface {
	CreateUser(c context.Context, dto DTO.CreateUserRequest) error
	FindUserByEmail(c context.Context, email string) (*domain.UserEntity, error)
}
