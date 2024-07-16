package auth

import "github.com/labstack/echo/v4"

type AuthService interface {
	Authenticate(c echo.Context, email, password string) (string, string, error)
	RefreshToken(c echo.Context, refreshToken string) (string, string, error)
	Generate2FACode(c echo.Context, username string) (string, error)
	Validate2FACode(c echo.Context, username, code string) error
	GenerateTokens(c echo.Context, email string) (string, string, error)
}
