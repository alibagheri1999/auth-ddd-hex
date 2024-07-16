package user

import (
	"DDD-HEX/internal/application/utils"
	"DDD-HEX/internal/domain"
	"DDD-HEX/internal/ports/repository"
	"errors"
	"github.com/labstack/echo/v4"
)

type userServiceImpl struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userServiceImpl{userRepository: userRepo}
}

func (s *userServiceImpl) CreateUser(c echo.Context, name, email, password string) error {
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

func (s *userServiceImpl) FindUserByEmail(c echo.Context, email string) (*domain.User, error) {
	if !isValidEmail(email) {
		return nil, errors.New("invalid email format")
	}
	user, err := s.userRepository.FindByEmail(email)
	return &user, err
}
