package cache

import (
	"context"
	"time"
)

// Pinger is an interface for checking the connection to the database
type Pinger interface {
	Ping(ctx context.Context) error
}

// Cache is an interface for caching
type Cache interface {
	HSetField(ctx context.Context, key string, field string, value interface{}) error
	HSet(ctx context.Context, key string, values interface{}) error
	Set(ctx context.Context, key string, value interface{}) error
	HGetAll(ctx context.Context, key string) (map[string]string, error)
	Get(ctx context.Context, key string) (string, error)
	Del(ctx context.Context, keys ...string) error
	Expire(ctx context.Context, key string, expiration time.Duration) error
	Exists(ctx context.Context, key string) (bool, error)
}

// Client is an interface for caching
type Client interface {
	Pinger
	Cache
}
