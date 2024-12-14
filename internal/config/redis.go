package config

import (
	"net"
	"os"

	"github.com/pkg/errors"
)

const (
	redisHostEnvName     = "REDIS_HOST"
	redisPortEnvName     = "REDIS_PORT"
	redisPasswordEnvName = "REDIS_PASSWORD"
)

// RedisConfig is an interface for redis configs.
type RedisConfig interface {
	Address() string
	Password() string
}

type redisConfig struct {
	host string
	port string

	password string
}

// NewRedisConfig creates a new redis config.
func NewRedisConfig() (RedisConfig, error) {
	host := os.Getenv(redisHostEnvName)
	if len(host) == 0 {
		return nil, errors.New("redis host not found")
	}

	port := os.Getenv(redisPortEnvName)
	if len(port) == 0 {
		return nil, errors.New("redis port not found")
	}

	password := os.Getenv(redisPasswordEnvName)
	if len(password) == 0 {
		return nil, errors.New("redis password not found")
	}

	return &redisConfig{
		host:     host,
		port:     port,
		password: password,
	}, nil
}

// Address returns redis address.
func (cfg *redisConfig) Address() string {
	return net.JoinHostPort(cfg.host, cfg.port)
}

func (cfg *redisConfig) Password() string {
	return cfg.password
}
