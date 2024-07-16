package setup

import (
	Handler "DDD-HEX/internal/adapters/http"
	"DDD-HEX/internal/application/services/auth"
	"DDD-HEX/internal/application/services/user"
)

type Handlers struct {
	UserHandler *Handler.UserHandler
	AuthHandler *Handler.AuthHandler
}

func NewHandler(userService user.UserService, authService auth.AuthService) *Handlers {
	userHandler := &Handler.UserHandler{UserService: userService}
	authHandler := &Handler.AuthHandler{AuthService: authService}
	return &Handlers{
		UserHandler: userHandler,
		AuthHandler: authHandler,
	}
}
