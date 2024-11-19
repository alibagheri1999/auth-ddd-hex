package cache

import (
	"DDD-HEX/internal/adapters/cache/redis"
	"DDD-HEX/internal/application/utils"
	"DDD-HEX/internal/ports/cache"
	"DDD-HEX/internal/ports/clients"
	"errors"
	"fmt"
)

type Repository interface {
	NewAuthCacheRepository(cache clients.Cache) cache.CacheRepository
}

type CacheRepoFactory struct {
	Cache clients.Cache
}

func (f *CacheRepoFactory) NewAuthCacheRepository(cache clients.Cache) cache.CacheRepository {
	return &redis.CacheRepository{Cache: cache}
}

func NewCacheRepository(cacheType string, cache clients.Cache) (Repository, error) {
	cacheEnum, err := utils.ParseCacheEnum(cacheType)
	if err != nil {
		return nil, fmt.Errorf("unsupported cache type: %s", cacheType)
	}
	switch cacheEnum {
	case utils.Redis:
		return &CacheRepoFactory{Cache: cache}, nil
	default:
		return nil, errors.New(fmt.Sprintf("unsupported cache type: %s", cacheType))
	}
}
