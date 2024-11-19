package main

import (
	"DDD-HEX/cmd/server"
	"DDD-HEX/cmd/setup"
	"DDD-HEX/internal/adapters/cache"
	"DDD-HEX/internal/adapters/db"
	"DDD-HEX/internal/application/utils"
	middlewares "DDD-HEX/pkg/middleware"
	"context"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"time"
)

func main() {
	utils.LoggerSetup()
	config := utils.ConfigSetup()
	appCfg := config.App
	dbCfg := config.Database
	cacheCfg := config.Cache
	// Redis context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cache, err := cache.NewCacheFactory(ctx, appCfg.CacheType, cacheCfg)
	if err != nil {
		logrus.Fatal("Failed to initialize cache:", err)
	}
	go func() {
		err := cache.EnsureConnected(3)
		if err != nil {
			logrus.Fatal("Failed to ensure Redis connection:", err)
		}
	}()
	DB, err := db.NewDatabase(appCfg, dbCfg)
	if err != nil {
		logrus.Fatal("Failed to connect to the database:", err)
	}
	repositories := setup.SetupRepositories(appCfg, DB, cache)
	services := setup.SetupServices(repositories, appCfg)
	handlers := setup.NewHandler(services)
	router := server.NewRouter()
	router.Use(middlewares.HealthCheck(DB))
	server.RegisterRoutes(router, handlers)
	server.NewServer(router, appCfg.Port, time.Duration(1)).StartListening()
}
