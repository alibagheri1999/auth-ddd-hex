package cmd

import (
	"DDD-HEX/internal/adapters/db/postgres"
	"DDD-HEX/internal/ports/repository"
	"database/sql"
)

func setupRepositories(db *sql.DB) (repository.UserRepository, repository.AuthRepository, repository.ProductRepository) {
	userRepository := &postgres.UserRepository{DB: db}
	authRepository := &postgres.AuthRepository{DB: db}
	productRepository := &postgres.ProductRepository{DB: db}

	return userRepository, authRepository, productRepository
}
