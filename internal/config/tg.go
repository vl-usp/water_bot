package config

import (
	"errors"
	"os"
)

const (
	tgBotTokenEnvName = "TELEGRAM_BOT_TOKEN"
)

// TGConfig is an interface that defines methods for Telegram Bot configuration.
type TGConfig interface {
	Token() string
}

// tgConfig contains the configuration for a Telegram Bot.
type tgConfig struct {
	token string
}

// NewTGConfig creates a new TGConfig by getting the Telegram Bot token
func NewTGConfig() (TGConfig, error) {
	token := os.Getenv(tgBotTokenEnvName)
	if len(token) == 0 {
		return nil, errors.New("telegram bot token is not set")
	}

	tgConf := &tgConfig{
		token: token,
	}

	return tgConf, nil
}

// Token returns the Telegram Bot token.
func (c *tgConfig) Token() string {
	return c.token
}
