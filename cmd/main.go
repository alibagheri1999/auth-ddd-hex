package main

import (
	"DDD-HEX/cmd/server"
	"DDD-HEX/cmd/setup"
	"DDD-HEX/internal/adapters/db/postgres"
	"DDD-HEX/internal/adapters/db/redis"
	"DDD-HEX/internal/application/utils"
	middlewares "DDD-HEX/pkg/middleware"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	"log"
	"time"
)

func main() {
	config := utils.ConfigSetup()
	appCfg := config.App
	dbCfg := config.Postgres
	cacheCfg := config.Redis
	cache := redis.NewRedisClient(cacheCfg)
	db, err := postgres.NewRepositories(appCfg, dbCfg)
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	userRepository, authRepository, cacheRepository := setup.SetupRepositories(db, cache)
	userService, authService := setup.SetupServices(userRepository, authRepository, cacheRepository, appCfg)
	handlers := setup.NewHandler(userService, authService)
	router := server.NewRouter()
	router.Use(middlewares.LoggingMiddleware)
	router.Use(middlewares.SecurityHeaders)
	router.Use(middleware.Recover())
	router.Use(middleware.CORS())
	router.Use(middlewares.CORS())
	router.Use(middlewares.HealthCheck(db))
	server.RegisterRoutes(router, handlers)
	server.NewServer(router, appCfg.Port, time.Duration(1)).StartListening()
}
