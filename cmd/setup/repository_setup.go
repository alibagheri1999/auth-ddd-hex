package setup

import (
	"DDD-HEX/internal/adapters/db/postgres"
	redisAdapter "DDD-HEX/internal/adapters/db/redis"
	"DDD-HEX/internal/ports/repository"
	"github.com/go-redis/redis/v8"
)

func SetupRepositories(db *postgres.Repositories, redisClient *redis.Client) (repository.UserRepository, repository.AuthRepository, repository.CacheRepository) {
	userRepository := &postgres.UserRepository{DB: db.DB()}
	authRepository := &postgres.AuthRepository{DB: db.DB()}
	cacheRepository := &redisAdapter.CacheRepository{RedisClient: redisClient}
	return userRepository, authRepository, cacheRepository
}
