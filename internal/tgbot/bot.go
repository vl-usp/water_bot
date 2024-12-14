package tgbot

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/fsm"
	"github.com/vl-usp/water_bot/internal/service"
	"github.com/vl-usp/water_bot/pkg/logger"
)

// TGBot represents a Telegram bot client.
type TGBot struct {
	userService service.UserService
	bot         *bot.Bot
	fsm         *fsm.FSM
}

// New creates a new Telegram bot client.
func New(token string, userService service.UserService) (*TGBot, error) {
	tgbot := &TGBot{
		userService: userService,
	}

	opts := []bot.Option{
		bot.WithDefaultHandler(tgbot.defaultHandler),
	}

	b, err := bot.New(token, opts...)
	if err != nil {
		return nil, err
	}

	tgbot.bot = b

	tgbot.setHandlers()
	tgbot.initFSM()

	return tgbot, nil
}

// Run starts the Telegram bot.
func (tgbot *TGBot) Run(ctx context.Context) {
	logger.Get("tgbot", "tgbot.Run").Info("running tgbot...")
	tgbot.bot.Start(ctx)
}

func (tgbot *TGBot) setHandlers() {
	tgbot.bot.RegisterHandler(bot.HandlerTypeMessageText, "/start", bot.MatchTypeExact, tgbot.startHandler)
}
