package db

import (
	"DDD-HEX/config"
	"DDD-HEX/internal/adapters/db/postgres"
	"DDD-HEX/internal/ports/clients"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
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
	logrus.Printf("Successfully connected to %s\n", appConfig.DbType)
	defer database.Close()
	return database, nil
}
