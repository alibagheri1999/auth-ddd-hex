package setup

import (
	"DDD-HEX/internal/adapters/db/postgres"
	redisAdapter "DDD-HEX/internal/adapters/db/redis"
	"DDD-HEX/internal/ports/cache"
	"DDD-HEX/internal/ports/repository"
)

func SetupRepositories(db *postgres.Repositories, redisClient *redisAdapter.ClientWrapper) (repository.UserRepository, repository.AuthRepository, cache.CacheRepository) {
	userRepository := &postgres.UserRepository{DB: db.DB()}
	authRepository := &postgres.AuthRepository{DB: db.DB()}
	cacheRepository := &redisAdapter.CacheRepository{RedisClient: redisClient}
	return userRepository, authRepository, cacheRepository
}
