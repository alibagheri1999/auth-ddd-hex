package auth

import (
	"context"
)

type AuthService interface {
	Authenticate(c context.Context, email, password string) (string, string, error)
	Validate2FA(c context.Context, email string) (string, string, error)
	RefreshToken(c context.Context, refreshToken string) (string, string, error)
	Generate2FACode(c context.Context, username string) (string, error)
	Validate2FACode(c context.Context, username, code string) error
}
