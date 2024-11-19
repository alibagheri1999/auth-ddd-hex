package cache

import (
	"DDD-HEX/config"
	"DDD-HEX/internal/adapters/cache/redis"
	"DDD-HEX/internal/ports/clients"
	"context"
	"fmt"
)

func NewCacheFactory(ctx context.Context, cacheType string, config config.CacheConfig) (clients.Cache, error) {
	switch cacheType {
	case "redis":
		return redis.NewRedisCache(ctx, config), nil
	default:
		return nil, fmt.Errorf("unsupported cache type: %s", cacheType)
	}
}
