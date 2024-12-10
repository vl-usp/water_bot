package main

import (
	"context"

	"github.com/vl-usp/water_bot/internal/app"
	"github.com/vl-usp/water_bot/pkg/logger"
)

func main() {
	ctx := context.Background()

	a, err := app.NewApp(ctx)
	if err != nil {
		logger.Get("main", "main").Error("failed to init app", "error", err.Error())
	}
	logger.Get("main", "main").Info("app created")

	err = a.Run(ctx)
	if err != nil {
		logger.Get("main", "main").Error("failed to run app", "error", err.Error())
	}
}
