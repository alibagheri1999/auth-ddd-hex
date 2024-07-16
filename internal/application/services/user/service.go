package user

import (
	"DDD-HEX/internal/domain"
	"github.com/labstack/echo/v4"
)

type UserService interface {
	CreateUser(c echo.Context, name, email, password string) error
	FindUserByEmail(c echo.Context, email string) (*domain.User, error)
}
