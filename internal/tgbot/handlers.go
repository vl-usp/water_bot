package tgbot

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/vl-usp/water_bot/internal/constants"
	"github.com/vl-usp/water_bot/internal/tgbot/converter"
	"github.com/vl-usp/water_bot/pkg/client/db/errors"
	"github.com/vl-usp/water_bot/pkg/logger"
)

func (client *Client) inputHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.Message == nil {
		return
	}

	userID := update.Message.From.ID

	switch client.fsm.Current(userID) {
	case constants.StateDefault:
		return
	case constants.StateOnboardingWeight:
		client.onboardingWeightHandler(ctx, b, update)
	default:
		logger.Get("tgbot", "tgbot.defaultHandler").Error("unexpected state", "state", client.fsm.Current(userID))
	}
}

func (client *Client) startHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	_, err := client.userService.CreateUser(ctx, *converter.ToUserFromTGUser(update.Message.From))
	if err != nil {
		if !errors.IsUniqueViolation(err) {
			logger.Get("tgbot", "tgbot.startHandler").Error("failed to create user", "error", err.Error())
		}
	}
	client.startOnboarding(ctx, b, update)
}
