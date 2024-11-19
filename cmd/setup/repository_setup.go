package setup

import (
	"DDD-HEX/config"
	cacheFactory "DDD-HEX/internal/adapters/cache"
	"DDD-HEX/internal/adapters/db"
	"DDD-HEX/internal/ports/cache"
	"DDD-HEX/internal/ports/clients"
	"DDD-HEX/internal/ports/repository"
	"log"
)

type Repositories struct {
	UserRepository  repository.UserRepository
	AuthRepository  repository.AuthRepository
	CacheRepository cache.CacheRepository
}

func SetupRepositories(appCfg config.AppConfig, DB clients.Database, cacheClient clients.Cache) *Repositories {
	repositories, err := db.NewRepository(appCfg.DbType, DB)
	if err != nil {
		log.Fatal("Failed to create repositories:", err)
	}
	cacheRepo, err := cacheFactory.NewCacheRepository(appCfg.CacheType, cacheClient)
	if err != nil {
		log.Fatal("Failed to initialize cache:", err)
	}
	return &Repositories{
		UserRepository:  repositories.NewUserRepository(DB),
		AuthRepository:  repositories.NewAuthRepository(DB),
		CacheRepository: cacheRepo.NewAuthCacheRepository(cacheClient),
	}
}
