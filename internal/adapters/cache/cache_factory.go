package cache

import (
	"DDD-HEX/config"
	"DDD-HEX/internal/adapters/cache/redis"
	"DDD-HEX/internal/ports/clients"
	"fmt"
)

func NewCacheFactory(cacheType string, config config.CacheConfig) (clients.Cache, error) {
	switch cacheType {
	case "redis":
		return redis.NewRedisCache(config), nil
	default:
		return nil, fmt.Errorf("unsupported cache type: %s", cacheType)
	}
}
