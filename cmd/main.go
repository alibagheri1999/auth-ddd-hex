package main

import (
	"DDD-HEX/cmd/server"
	"DDD-HEX/cmd/setup"
	"DDD-HEX/internal/adapters/cache/redis"
	"DDD-HEX/internal/adapters/db"
	"DDD-HEX/internal/application/utils"
	middlewares "DDD-HEX/pkg/middleware"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	"log"
	"time"
)

func main() {
	utils.LoggerSetup()
	config := utils.ConfigSetup()
	appCfg := config.App
	dbCfg := config.Database
	cacheCfg := config.Cache
	cache := redis.NewRedisClientWrapper(cacheCfg)
	go cache.EnsureConnected(3)
	DB, err := db.NewDatabase(appCfg, dbCfg)
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	userRepository, authRepository, cacheRepository := setup.SetupRepositories(appCfg, DB, cache)
	userService, authService := setup.SetupServices(userRepository, authRepository, cacheRepository, appCfg)
	handlers := setup.NewHandler(userService, authService)
	router := server.NewRouter()
	router.Use(middlewares.LoggingMiddleware)
	router.Use(middlewares.SecurityHeaders)
	router.Use(middleware.Recover())
	router.Use(middleware.CORS())
	router.Use(middlewares.CORS())
	router.Use(middlewares.HealthCheck(DB))
	//router.Use(middleware.Logger())
	server.RegisterRoutes(router, handlers)
	server.NewServer(router, appCfg.Port, time.Duration(1)).StartListening()
}
