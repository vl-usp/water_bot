package config

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/vl-usp/water_bot/internal/config/env"
)

// Load loads environment variables from a .env file at the given path.
func Load(path string) error {
	// Load the .env file
	err := godotenv.Load(path)
	if err != nil {
		// If loading the .env file fails, return a wrapped error
		return fmt.Errorf("failed to load config from path %s: %w", path, err)
	}

	return nil
}

// PGConfig is an interface that defines methods for PostgreSQL configuration.
type PGConfig interface {
	DSN() string
}

// TGBotConfig is an interface that defines methods for Telegram Bot configuration.
type TGBotConfig interface {
	Token() string
}

// LogConfig is an interface that defines methods for log configuration.
type LogConfig interface {
	DirPath() string
	Env() string
}

// Config contains all the configuration options.
type Config struct {
	PG    PGConfig
	TGBot TGBotConfig
	Log   LogConfig
}

// NewEnv creates a new Config by calling the respective New* functions from the env package and returns it.
// If any of the New* functions returns an error, it is wrapped into a fmt.Errorf and returned.
func NewEnv() (*Config, error) {
	// Get the PostgreSQL configuration
	pgConfig, err := env.NewPGConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to get pg config: %w", err)
	}

	// Get the Telegram Bot configuration
	tgBotConfig, err := env.NewTGBotConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to get tg bot config: %w", err)
	}

	// Get the log configuration
	logConfig, err := env.NewLogConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to get log config: %w", err)
	}

	// Create a new Config and return it
	return &Config{
		PG:    pgConfig,
		TGBot: tgBotConfig,
		Log:   logConfig,
	}, nil
}
