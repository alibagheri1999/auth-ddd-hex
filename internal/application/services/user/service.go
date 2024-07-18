package user

import (
	"DDD-HEX/internal/domain"
	"DDD-HEX/internal/domain/DTO"
	"github.com/labstack/echo/v4"
)

type UserService interface {
	CreateUser(c echo.Context, dto DTO.CreateUserRequest) error
	FindUserByEmail(c echo.Context, email string) (*domain.UserEntity, error)
}
