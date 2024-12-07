package bot

import (
	"context"
	"log/slog"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

// Bot wraps the go-telegram bot.Bot type.
type Bot struct {
	*bot.Bot
}

// New creates a new Telegram bot client.
func New(token string) (*Bot, error) {
	// Options for the bot.
	opts := []bot.Option{
		// Set the default handler for the bot.
		bot.WithDefaultHandler(handler),
	}

	// Create the bot client.
	b, err := bot.New(token, opts...)
	if err != nil {
		return nil, err
	}

	// Create a new Bot instance.
	return &Bot{
		Bot: b,
	}, nil
}

// handler is the default handler for the bot.
// It echoes the message that was received.
func handler(ctx context.Context, b *bot.Bot, update *models.Update) {
	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   update.Message.Text,
	})
	if err != nil {
		slog.Error("failed to send message", "error", err.Error())
	}
}
