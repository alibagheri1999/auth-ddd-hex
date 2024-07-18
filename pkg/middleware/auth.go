package middleware

import (
	"DDD-HEX/internal/application/services/auth"
	"DDD-HEX/internal/domain/DTO"
	"github.com/labstack/echo/v4"
	"net/http"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var res DTO.LoginResponse
		token, err := c.Request().Cookie("access_token")
		if err != nil {
			res.Message = "Forbidden"
			return c.JSON(http.StatusUnauthorized, res)
		}
		if token.Value == "" {
			res.Message = "Forbidden"
			return c.JSON(http.StatusForbidden, res)
		}
		claims, err := auth.ValidateToken(token.Value)
		if err != nil {
			res.Message = "Unauthorized"
			return c.JSON(http.StatusUnauthorized, res)
		}

		c.Set("user", claims)
		return next(c)
	}
}
