package tgbot

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/fsm"
	"github.com/go-telegram/ui/keyboard/inline"
	"github.com/vl-usp/water_bot/internal/model"
	"github.com/vl-usp/water_bot/pkg/logger"
)

// UserService provide user data.
type UserService interface {
	CreateUser(ctx context.Context, user model.User) error
	UpdateUserFromCache(ctx context.Context, userID int64) error
	GetUser(ctx context.Context, userID int64) (*model.User, error)
	GetFullUser(ctx context.Context, userID int64) (*model.User, error)
	SaveUserParam(ctx context.Context, userID int64, field string, value interface{}) error
}

// ReferenceProvider provide reference data.
type ReferenceProvider interface {
	SexList(ctx context.Context) ([]model.Sex, error)
	PhysicalActivityList(ctx context.Context) ([]model.PhysicalActivity, error)
	ClimateList(ctx context.Context) ([]model.Climate, error)
	TimezoneList(ctx context.Context) ([]model.Timezone, error)
}

// Client represents a Telegram bot client.
type Client struct {
	userService UserService
	refProvider ReferenceProvider
	bot         *bot.Bot
	fsm         *fsm.FSM
}

// New creates a new Telegram bot client.
func New(token string, userService UserService, refProvider ReferenceProvider) (*Client, error) {
	client := &Client{
		userService: userService,
		refProvider: refProvider,
	}

	opts := []bot.Option{
		bot.WithDefaultHandler(client.inputHandler),
	}

	b, err := bot.New(token, opts...)
	if err != nil {
		return nil, err
	}

	client.bot = b

	client.setHandlers()

	return client, nil
}

// Run starts the Telegram bot.
func (client *Client) Run(ctx context.Context) {
	logger.Get("tgbot", "client.Run").Info("running tgbot...")
	client.bot.Start(ctx)
}

func (client *Client) sendMessage(ctx context.Context, chatID int64, text string) {
	_, err := client.bot.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: chatID,
		Text:   text,
	})
	if err != nil {
		logger.Get("tgbot", "client.sendMessage").Error("failed to send message", "error", err.Error())
		return
	}
}

func (client *Client) sendMessageWithKeyboard(ctx context.Context, chatID int64, text string, keyboard *inline.Keyboard) {
	_, err := client.bot.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      chatID,
		Text:        text,
		ReplyMarkup: keyboard,
	})
	if err != nil {
		logger.Get("tgbot", "client.sendMessageWithKeyboard").Error("failed to send message", "error", err.Error())
		return
	}
}

func (client *Client) sendErrorMessage(ctx context.Context, chatID int64, text string) {
	_, err := client.bot.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: chatID,
		Text:   text,
	})
	if err != nil {
		logger.Get("tgbot", "client.sendErrorMessage").Error("failed to send message", "error", err.Error())
		return
	}
}

func (client *Client) setHandlers() {
	client.bot.RegisterHandler(bot.HandlerTypeMessageText, "/start", bot.MatchTypeExact, client.startHandler)
}
