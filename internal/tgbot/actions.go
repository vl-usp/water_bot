package tgbot

import (
	"context"
	"fmt"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/vl-usp/water_bot/internal/client/db/errors"
	"github.com/vl-usp/water_bot/internal/tgbot/converter"
)

func (tgbot *TGBot) createUserAction(ctx context.Context, b *bot.Bot, update *models.Update) error {
	id, err := tgbot.userService.Create(ctx, converter.ToUserFromTGUser(update.Message.From))
	if err != nil {
		if errors.IsUniqueViolation(err) {
			_, err = b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: update.Message.Chat.ID,
				Text:   fmt.Sprintf("Hello, user #%d! You are already registered!", update.Message.From.ID),
			})
			return err
		}

		return err
	}

	_, err = b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   fmt.Sprintf("Hello, user #%d!", id),
	})
	return err
}
