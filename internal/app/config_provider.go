package app

import (
	"github.com/vl-usp/water_bot/internal/config"
	"github.com/vl-usp/water_bot/pkg/client/cache/redis"
	"github.com/vl-usp/water_bot/pkg/logger"
)

type configProvider struct {
	systemConfig config.SystemConfig
	pgConfig     config.PGConfig
	tgConfig     config.TGConfig
	redisConfig  config.RedisConfig
}

func newConfigProvider() *configProvider {
	return &configProvider{}
}

// SystemConfig returns a config that stores system settings.
func (c *configProvider) SystemConfig() config.SystemConfig {
	if c.systemConfig == nil {
		cfg, err := config.NewSystemConfig()
		if err != nil {
			logger.Get("app", "s.SystemConfig").Error("failed to get system config", "error", err.Error())
		}

		c.systemConfig = cfg
	}

	return c.systemConfig
}

// PGConfig returns a pg config.
func (c *configProvider) PGConfig() config.PGConfig {
	if c.pgConfig == nil {
		cfg, err := config.NewPGConfig()
		if err != nil {
			logger.Get("app", "s.PGConfig").Error("failed to get pg config", "error", err.Error())
		}

		c.pgConfig = cfg
	}

	return c.pgConfig
}

// TGConfig returns a tg config.
func (c *configProvider) TGConfig() config.TGConfig {
	if c.tgConfig == nil {
		cfg, err := config.NewTGConfig()
		if err != nil {
			logger.Get("app", "s.TGConfig").Error("failed to get tg config", "error", err.Error())
		}

		c.tgConfig = cfg
	}

	return c.tgConfig
}

// RedisConfig returns a redis config.
func (c *configProvider) RedisConfig() redis.RedisConfig {
	if c.redisConfig == nil {
		cfg, err := config.NewRedisConfig()
		if err != nil {
			logger.Get("app", "s.RedisConfig").Error("failed to get redis config", "error", err.Error())
		}

		c.redisConfig = cfg
	}

	return c.redisConfig
}
