package config

import (
	"log/slog"
	"sync"

	"github.com/cristalhq/aconfig"
	"github.com/cristalhq/aconfig/aconfigyaml"
)

type Config struct {
	TelegramBotToken string `yaml:"telegram_bot_token" env:"TELEGRAM_BOT_TOKEN" required:"true"`
}

var (
	cfg  Config
	once sync.Once
)

func Get() Config {
	once.Do(func() {
		loader := aconfig.LoaderFor(&cfg, aconfig.Config{
			EnvPrefix: "WB",
			Files:     []string{"./config.yaml", "./config.local.yaml", "$HOME/.config/water_bot/config.yaml"},
			FileDecoders: map[string]aconfig.FileDecoder{
				".yaml": aconfigyaml.New(),
			},
		})

		if err := loader.Load(); err != nil {
			slog.Error("failed to load config", "error", err)
		}
	})

	return cfg
}
