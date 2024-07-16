package server

import (
	"DDD-HEX/cmd/setup"
	"github.com/labstack/echo/v4"
	_ "github.com/swaggo/swag"
)

// RegisterRoutes register all routes that we want to use in app
func RegisterRoutes(router *echo.Echo, handler *setup.Handlers) {

	v1 := router.Group("/api/v1")

	//v1.GET("/health-check", handler.GeneralService.CheckHealth)

	v1.POST("/users", handler.UserHandler.CreateUser)
	v1.POST("/login", handler.AuthHandler.Login)
	v1.POST("/refresh", handler.AuthHandler.Refresh)
	v1.POST("/generate-2FA-code", handler.AuthHandler.Generate2FACode)
	v1.POST("/validate-2FA-code", handler.AuthHandler.Validate2FACode)
}
