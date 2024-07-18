package repository

import "DDD-HEX/internal/domain"

type AuthRepository interface {
	Save(auth domain.AuthEntity) error
	FindByAccessToken(token string) (domain.AuthEntity, error)
	FindByRefreshToken(token string) (domain.UserAuthEntity, error)
}
