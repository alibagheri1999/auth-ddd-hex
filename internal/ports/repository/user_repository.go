package repository

import "DDD-HEX/internal/domain"

type UserRepository interface {
	Save(user domain.UserEntity) error
	FindByID(id string) (domain.UserEntity, error)
	FindByEmail(email string) (domain.UserEntity, error)
}
