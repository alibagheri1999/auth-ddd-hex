package setup

import (
	Handler "DDD-HEX/internal/adapters/http"
)

type Handlers struct {
	UserHandler *Handler.UserHandler
	AuthHandler *Handler.AuthHandler
}

func NewHandler(services *Services) *Handlers {
	userHandler := &Handler.UserHandler{UserService: *services.UserService}
	authHandler := &Handler.AuthHandler{AuthService: *services.AuthService}
	return &Handlers{
		UserHandler: userHandler,
		AuthHandler: authHandler,
	}
}
