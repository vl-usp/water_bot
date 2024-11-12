package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/vl-usp/water_bot/internal/bot"
	"github.com/vl-usp/water_bot/internal/config"
	"github.com/vl-usp/water_bot/pkg/logger"
)

func main() {
	log := logger.Get()
	defer logger.CloseFile()

	cfg := config.Get()
	log.Info("config loaded")

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	bot, err := bot.New(cfg.TelegramBotToken)
	if err != nil {
		log.Error("failed to create bot", "error", err)
		panic(err)
	}
	log.Info("bot created")

	log.Info("starting the bot")
	bot.Start(ctx)
}
