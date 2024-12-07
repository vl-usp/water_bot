package env

import (
	"errors"
	"os"
)

const (
	tgBotTokenEnvName = "TELEGRAM_BOT_TOKEN"
)

// TGBotConfig contains the configuration for a Telegram Bot.
type TGBotConfig struct {
	token string
}

// NewTGBotConfig creates a new TGBotConfig by getting the Telegram Bot token
func NewTGBotConfig() (*TGBotConfig, error) {
	token := os.Getenv(tgBotTokenEnvName)
	if len(token) == 0 {
		return nil, errors.New("telegram bot token is not set")
	}

	return &TGBotConfig{
		token: token,
	}, nil
}

// Token returns the Telegram Bot token.
func (c *TGBotConfig) Token() string {
	return c.token
}
