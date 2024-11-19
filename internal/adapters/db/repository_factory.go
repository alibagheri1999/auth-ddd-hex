package db

import (
	"DDD-HEX/config"
	"DDD-HEX/internal/adapters/db/postgres"
	"DDD-HEX/internal/ports/clients"
	"errors"
	"fmt"
)

type PostgresRepository struct {
	DB clients.Database
}

type Repository interface {
	NewAuthRepository(database clients.Database) *postgres.AuthRepository
	NewUserRepository(database clients.Database) *postgres.UserRepository
}

// NewUserRepository creates a new UserRepository.
func (p *PostgresRepository) NewUserRepository(database clients.Database) *postgres.UserRepository {
	return &postgres.UserRepository{DB: database}
}

// NewAuthRepository creates a new AuthRepository.
func (p *PostgresRepository) NewAuthRepository(database clients.Database) *postgres.AuthRepository {
	return &postgres.AuthRepository{DB: database}
}

func NewRepository(appConfig config.AppConfig, database clients.Database) (Repository, error) {
	switch appConfig.DbType {
	case "postgres":
		return &PostgresRepository{DB: database}, nil
	default:
		return nil, errors.New(fmt.Sprintf("unsupported database type: %s", appConfig.DbType))
	}
}
