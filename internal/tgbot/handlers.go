package tgbot

import (
	"context"
	"fmt"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/vl-usp/water_bot/internal/client/db/errors"
	"github.com/vl-usp/water_bot/internal/model"
	"github.com/vl-usp/water_bot/pkg/logger"
)

func (tgbot *TGBot) startHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	logger.Get("tgbot", "tgbot.startHandler").Info("creating user...", "user", update.Message.From)

	id, err := tgbot.userService.Create(ctx, &model.User{
		ID:           update.Message.From.ID,
		FirstName:    update.Message.From.FirstName,
		LastName:     update.Message.From.LastName,
		Username:     update.Message.From.Username,
		LanguageCode: update.Message.From.LanguageCode,
	})
	if err != nil {
		if errors.IsUniqueViolation(err) {
			_, err = b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: update.Message.Chat.ID,
				Text:   fmt.Sprintf("Hello, user #%d! You are already registered", id),
			})
			if err != nil {
				logger.Get("tgbot", "tgbot.startHandler").Error("failed to send message", "error", err.Error())
			}
			return
		}

		logger.Get("tgbot", "tgbot.startHandler").Error("failed to create user", "error", err.Error())

	}
	// init onboarding

	_, err = b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   fmt.Sprintf("Hello, user #%d!", id),
	})
	if err != nil {
		logger.Get("tgbot", "tgbot.startHandler").Error("failed to send message", "error", err.Error())
	}
}
