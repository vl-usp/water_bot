package tgbot

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/vl-usp/water_bot/pkg/logger"
)

func (tgbot *TGBot) startHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	err := tgbot.createUserAction(ctx, b, update)
	if err != nil {
		logger.Get("tgbot", "tgbot.startHandler").Error("failed to create user", "error", err.Error())
	}
	// init onboarding

	// tgbot.onboardingHandler(ctx, b, update)
}
