package redis

import (
	"context"
	"log/slog"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/vl-usp/water_bot/pkg/client/cache"
)

type handler func(ctx context.Context, conn redis.Conn) error

type redisClient struct {
	rdb *redis.Client
	log *slog.Logger
}

// New creates a new redis client.
func New(cfg RedisConfig, log *slog.Logger) cache.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Address(),
		Password: cfg.Password(),
		DB:       0,
	})

	return &redisClient{
		rdb: rdb,
		log: log,
	}
}

func (c *redisClient) HSetField(ctx context.Context, key string, field string, value interface{}) error {
	return c.rdb.HSet(ctx, key, field, value).Err()
}

// HashSet sets hash values.
func (c *redisClient) HSet(ctx context.Context, key string, values interface{}) error {
	return c.rdb.HSet(ctx, key, values).Err()
}

// Set sets value.
func (c *redisClient) Set(ctx context.Context, key string, value interface{}) error {
	return c.rdb.Set(ctx, key, value, 0).Err()
}

// HGetAll returns all hash values.
func (c *redisClient) HGetAll(ctx context.Context, key string) (map[string]string, error) {
	return c.rdb.HGetAll(ctx, key).Result()
}

// Get returns value.
func (c *redisClient) Get(ctx context.Context, key string) (string, error) {
	return c.rdb.Get(ctx, key).Result()
}

func (c *redisClient) Del(ctx context.Context, keys ...string) error {
	return c.rdb.Del(ctx, keys...).Err()
}

// Expire expires key.
func (c *redisClient) Expire(ctx context.Context, key string, expiration time.Duration) error {
	return c.rdb.Expire(ctx, key, expiration).Err()
}

func (c *redisClient) Exists(ctx context.Context, key string) (bool, error) {
	v, err := c.rdb.Exists(ctx, key).Result()
	if err != nil {
		return false, err
	}

	return v != 0, nil
}

// Ping pings the database.
func (c *redisClient) Ping(ctx context.Context) error {
	return c.rdb.Ping(ctx).Err()
}
