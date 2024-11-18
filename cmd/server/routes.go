package server

import (
	"DDD-HEX/cmd/setup"
	"DDD-HEX/pkg/middleware"
	"github.com/labstack/echo/v4"
	_ "github.com/swaggo/swag"
)

// RegisterRoutes register all routes that we want to use in app
func RegisterRoutes(router *echo.Echo, handler *setup.Handlers) {

	v1 := router.Group("/api/v1")
	auth := v1.Group("/auth")
	auth.POST("/login", handler.AuthHandler.Login)
	auth.POST("/refresh", handler.AuthHandler.Refresh)
	auth.POST("/generate-2FA-code", handler.AuthHandler.Generate2FACode)
	auth.POST("/validate-2FA-code", handler.AuthHandler.Validate2FACode)
	auth.POST("/users", handler.UserHandler.CreateUser)
	protected := v1.Group("/core")
	protected.Use(middleware.AuthMiddleware)

	admin := protected.Group("/admin")
	admin.Use(middleware.RBACMiddleware("admin"))
}
