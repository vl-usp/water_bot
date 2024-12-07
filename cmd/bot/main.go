package main

import (
	"context"
	"flag"
	"log/slog"
	"os"
	"os/signal"

	"github.com/vl-usp/water_bot/internal/bot"
	"github.com/vl-usp/water_bot/internal/config"
	"github.com/vl-usp/water_bot/pkg/logger"
)

var envPath string

func init() {
	flag.StringVar(&envPath, "env", ".env", "path to .env file")
}

func main() {
	flag.Parse()
	if err := config.Load(envPath); err != nil {
		slog.Error("failed to load config", "error", err.Error())
		panic("failed to load config: " + err.Error())
	}

	cfg, err := config.NewEnv()
	if err != nil {
		slog.Error("failed to create config", "error", err.Error())
		panic("failed to create config: " + err.Error())
	}

	log, file := logger.SetupLogger(cfg.Log.Env(), cfg.Log.DirPath())
	defer func() {
		err := file.Close()
		if err != nil {
			log.Error("failed to close log file", "error", err.Error())
		}
	}()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	bot, err := bot.New(cfg.TGBot.Token())
	if err != nil {
		log.Error("failed to create bot", "error", err)
		panic(err)
	}
	log.Info("bot created")

	log.Info("starting the bot")
	bot.Start(ctx)
}
