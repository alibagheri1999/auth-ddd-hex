package middleware

import (
	"DDD-HEX/internal/application/services/auth"
	"DDD-HEX/internal/domain/DTO"
	"github.com/labstack/echo/v4"
	"net/http"
)

func RBACMiddleware(roles ...string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			var res DTO.LoginResponse
			user := c.Get("user").(*auth.Claims)
			for _, role := range roles {
				if user.Role == role {
					return next(c)
				}
			}
			res.Message = "Forbidden"
			return c.JSON(http.StatusForbidden, res)
		}
	}
}
