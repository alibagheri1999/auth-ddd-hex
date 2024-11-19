package setup

import (
	"DDD-HEX/config"
	"DDD-HEX/internal/adapters/cache/redis"
	"DDD-HEX/internal/adapters/db"
	"DDD-HEX/internal/ports/cache"
	"DDD-HEX/internal/ports/clients"
	"DDD-HEX/internal/ports/repository"
	"log"
)

func SetupRepositories(appConfig config.AppConfig, DB clients.Database, redisClient *redis.ClientWrapper) (repository.UserRepository, repository.AuthRepository, cache.CacheRepository) {
	repositories, err := db.NewRepository(appConfig, DB)
	if err != nil {
		log.Fatal("Failed to create repositories:", err)
	}
	cacheRepository := &redis.CacheRepository{RedisClient: redisClient}
	return repositories.NewUserRepository(DB), repositories.NewAuthRepository(DB), cacheRepository
}
