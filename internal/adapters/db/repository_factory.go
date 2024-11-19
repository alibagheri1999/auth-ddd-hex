package db

import (
	"DDD-HEX/internal/adapters/db/postgres"
	"DDD-HEX/internal/application/utils"
	"DDD-HEX/internal/ports/clients"
	"DDD-HEX/internal/ports/repository"
	"errors"
	"fmt"
)

type DbRepository struct {
	DB clients.Database
}

type Repository interface {
	NewAuthRepository(database clients.Database) repository.AuthRepository
	NewUserRepository(database clients.Database) repository.UserRepository
}

// NewUserRepository creates a new UserRepository.
func (p *DbRepository) NewUserRepository(database clients.Database) repository.UserRepository {
	return &postgres.UserRepository{DB: database}
}

// NewAuthRepository creates a new AuthRepository.
func (p *DbRepository) NewAuthRepository(database clients.Database) repository.AuthRepository {
	return &postgres.AuthRepository{DB: database}
}

func NewRepository(dbType string, database clients.Database) (Repository, error) {
	dbEnum, err := utils.ParseDbEnum(dbType)
	if err != nil {
		return nil, fmt.Errorf("unsupported db type: %s", dbType)
	}
	switch dbEnum {
	case utils.Postgres:
		return &DbRepository{DB: database}, nil
	default:
		return nil, errors.New(fmt.Sprintf("unsupported database type: %s", dbType))
	}
}
