package repository

import "DDD-HEX/internal/domain"

type UserRepository interface {
	Save(user domain.User) error
	FindByID(id string) (domain.User, error)
	FindByEmail(email string) (domain.User, error)
}
