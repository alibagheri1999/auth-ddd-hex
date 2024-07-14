package user

import (
	"DDD-HEX/internal/application/utils"
	"DDD-HEX/internal/domain"
	"DDD-HEX/internal/ports/repository"
	"errors"
)

type userServiceImpl struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userServiceImpl{userRepository: userRepo}
}

func (s *userServiceImpl) CreateUser(name, email, password string) error {
	if !isValidEmail(email) {
		return errors.New("invalid email format")
	}
	user := domain.User{
		Name:     name,
		Email:    email,
		Password: utils.Hash(password), // Assume this function hashes the password
	}
	return s.userRepository.Save(user)
}
