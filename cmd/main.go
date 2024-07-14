package cmd

import (
	"DDD-HEX/internal/application/utils"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
)

func main() {
	config := utils.ConfigSetup()
	appCfg := config.App
	dbCfg := config.Postgres
	db, err := sql.Open(appCfg.DbName, utils.GeneratePostgresConnectionString(dbCfg))
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	userRepository, authRepository, productRepository := setupRepositories(db)
	userService, authService, productService := setupServices(userRepository, authRepository, productRepository, appCfg)
	router := setupHandlers(userService, authService, productService)

	logrus.Info("Starting server on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
