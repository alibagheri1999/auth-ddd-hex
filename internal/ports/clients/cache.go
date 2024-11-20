package clients

import "context"

type Cache interface {
	Connect() error
	Close() error
	Ping(ctx context.Context) error
	EnsureConnected(maxRetries int) error
	Set(ctx context.Context, key string, value interface{}, ttlSeconds uint32) error
	Get(ctx context.Context, key string) (string, error)
	Delete(ctx context.Context, key string) error
	Flush(ctx context.Context) error
}
