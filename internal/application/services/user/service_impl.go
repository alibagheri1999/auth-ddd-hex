package user

import (
	"DDD-HEX/internal/application/utils"
	"DDD-HEX/internal/domain"
	"DDD-HEX/internal/domain/DTO"
	"DDD-HEX/internal/ports/repository"
	"context"
	"database/sql"
	"errors"
)

type userServiceImpl struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userServiceImpl{userRepository: userRepo}
}

func (s *userServiceImpl) CreateUser(c context.Context, dto DTO.CreateUserRequest) error {
	if !isValidEmail(dto.Email) {
		return errors.New("invalid email format")
	}
	user := domain.UserEntity{
		Name:        dto.Name,
		Email:       dto.Email,
		Password:    sql.NullString{Valid: true, String: utils.Hash(dto.Password)},
		PhoneNumber: sql.NullString{Valid: true, String: dto.PhoneNumber},
	}
	return s.userRepository.Save(c, user)
}

func (s *userServiceImpl) FindUserByEmail(c context.Context, email string) (*domain.UserEntity, error) {
	if !isValidEmail(email) {
		return nil, errors.New("invalid email format")
	}
	user, err := s.userRepository.FindByEmail(c, email)
	return &user, err
}
