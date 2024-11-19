package db

import (
	"DDD-HEX/config"
	"DDD-HEX/internal/adapters/db/postgres"
	"DDD-HEX/internal/ports/clients"
	"errors"
	"fmt"
)

func NewDatabase(appConfig config.AppConfig, dbConfig config.DatabaseConfig) (clients.Database, error) {
	var database clients.Database
	switch appConfig.DbType {
	case "postgres":
		database = postgres.NewPostgresDB(dbConfig)
	default:
		return nil, errors.New(fmt.Sprintf("unsupported database type: %s", appConfig.DbType))
	}
	err := database.Connect()
	if err != nil {
		return nil, err
	}
	return database, nil
}
