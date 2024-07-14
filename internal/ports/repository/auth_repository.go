package repository

import "DDD-HEX/internal/domain"

type AuthRepository interface {
	Save(auth domain.Auth) error
	FindByAccessToken(token string) (domain.Auth, error)
	FindByRefreshToken(token string) (domain.Auth, error)
}
