package tgbot

import (
	"context"
	"strconv"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/go-telegram/fsm"
	"github.com/vl-usp/water_bot/internal/constants"
	"github.com/vl-usp/water_bot/internal/tgbot/converter"
	"github.com/vl-usp/water_bot/pkg/client/db/errors"
	"github.com/vl-usp/water_bot/pkg/logger"
)

const (
	stateDefault             fsm.StateID = "default"
	stateOnboardingAskWeight fsm.StateID = "onboarding_ask_weight"
)

func (tgbot *TGBot) initFSM() {
	tgbot.fsm = fsm.New(
		stateDefault,
		map[fsm.StateID]fsm.Callback{
			stateOnboardingAskWeight: tgbot.callbackAskWeight,
		},
	)
}

func (tgbot *TGBot) defaultHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.Message == nil {
		return
	}

	userID := update.Message.From.ID
	chatID := update.Message.Chat.ID

	switch tgbot.fsm.Current(userID) {
	case stateDefault:
		return
	case stateOnboardingAskWeight:
		weight, errWeight := strconv.Atoi(update.Message.Text)
		if errWeight != nil {
			_, err := b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: chatID,
				Text:   constants.WeightErrorMessage,
			})
			if err != nil {
				logger.Get("tgbot", "tgbot.defaultHandler").Error("failed to send message", "error", err.Error())
			}
			return
		}
		if weight < 30 || weight > 300 {
			_, err := b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: chatID,
				Text:   constants.WeightNotValidMessage,
			})
			if err != nil {
				logger.Get("tgbot", "tgbot.defaultHandler").Error("failed to send message", "error", err.Error())
			}
			return
		}

		// Store the weight to cache
		err := tgbot.userService.SaveUserDataField(ctx, userID, constants.WeightKey, weight)
		if err != nil {
			logger.Get("tgbot", "tgbot.defaultHandler").Error("failed to save weight", "error", err.Error())
			return
		}

		tgbot.fsm.Transition(userID, stateDefault)
		tgbot.sexSelectHandler(ctx, b, update)
	default:
		logger.Get("tgbot", "tgbot.defaultHandler").Error("unexpected state", "state", tgbot.fsm.Current(userID))
	}
}

func (tgbot *TGBot) startHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	_, err := tgbot.userService.CreateUser(ctx, converter.ToUserFromTGUser(update.Message.From))
	if err != nil {
		if !errors.IsUniqueViolation(err) {
			logger.Get("tgbot", "tgbot.startHandler").Error("failed to create user", "error", err.Error())
		}
	}
	tgbot.onboardingHandler(ctx, b, update)
}
