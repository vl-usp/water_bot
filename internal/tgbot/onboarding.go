package tgbot

import (
	"context"
	"fmt"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/go-telegram/fsm"
	"github.com/go-telegram/ui/keyboard/inline"
	"github.com/vl-usp/water_bot/internal/constants"
	"github.com/vl-usp/water_bot/pkg/logger"
)

func (tgbot *TGBot) onboardingHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   constants.BenifitOfWater,
	})
	if err != nil {
		logger.Get("tgbot", "tgbot.onboardingHandler").Error("failed to send message", "error", err.Error())
		return
	}

	_, err = b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   constants.StandardNorm,
	})
	if err != nil {
		logger.Get("tgbot", "tgbot.onboardingHandler").Error("failed to send message", "error", err.Error())
		return
	}

	keyboard := inline.New(b, inline.NoDeleteAfterClick()).
		Button(constants.WaterNormSelf, []byte(constants.WaterNormSelfKey), tgbot.onWaterNormSelect).
		Button(constants.WaterNormDefault, []byte(constants.WaterNormDefaultKey), tgbot.onWaterNormSelect)

	_, err = b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        constants.WaterNormQuestion,
		ReplyMarkup: keyboard,
	})
	if err != nil {
		logger.Get("tgbot", "tgbot.onboardingHandler").Error("failed to send message", "error", err.Error())
		return
	}
}

func (tgbot *TGBot) onWaterNormSelect(ctx context.Context, _ *bot.Bot, mes models.MaybeInaccessibleMessage, data []byte) {
	chatID := mes.Message.Chat.ID

	switch string(data) {
	case constants.WaterNormSelfKey:
		tgbot.fsm.Transition(chatID, stateOnboardingAskWeight, chatID)
	case constants.WaterNormDefaultKey:
		err := tgbot.userService.SaveUserDataField(ctx, chatID, constants.WaterGoalKey, constants.WaterGoalDefault)
		if err != nil {
			logger.Get("tgbot", "tgbot.onWaterNormSelect").Error("failed to save user data", "error", err.Error())
			return
		}

		userData, err := tgbot.userService.CreateUserData(ctx, chatID)
		if err != nil {
			logger.Get("tgbot", "tgbot.onWaterNormSelect").Error("failed to create user data", "error", err.Error())
			return
		}

		_, err = tgbot.bot.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: chatID,
			Text:   fmt.Sprintf(constants.WaterNormResult, userData.WaterGoal),
		})
		if err != nil {
			logger.Get("tgbot", "tgbot.onWaterNormSelect").Error("failed to send message", "error", err.Error())
			return
		}

	}
}

func (tgbot *TGBot) callbackAskWeight(_ *fsm.FSM, args ...any) {
	chatID := args[0]

	_, err := tgbot.bot.SendMessage(context.Background(), &bot.SendMessageParams{
		ChatID: chatID,
		Text:   constants.WeightQuestion,
	})
	if err != nil {
		logger.Get("tgbot", "tgbot.callbackAskWeight").Error("failed to send message", "error", err.Error())
	}
}

func (tgbot *TGBot) sexSelectHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	keyboard := inline.New(b, inline.NoDeleteAfterClick()).
		Button(constants.SexMale, []byte(constants.SexMaleKey), tgbot.onSexSelect).
		Button(constants.SexFemale, []byte(constants.SexFemaleKey), tgbot.onSexSelect)

	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        constants.SexQuestion,
		ReplyMarkup: keyboard,
	})
	if err != nil {
		logger.Get("tgbot", "tgbot.sexSelectHandler").Error("failed to send message", "error", err.Error())
		return
	}
}

func (tgbot *TGBot) onSexSelect(ctx context.Context, b *bot.Bot, mes models.MaybeInaccessibleMessage, data []byte) {
	chatID := mes.Message.Chat.ID
	err := tgbot.userService.SaveUserDataField(ctx, chatID, constants.SexKey, string(data))
	if err != nil {
		logger.Get("tgbot", "tgbot.onSexSelect").Error("failed to save user data", "error", err.Error())
		return
	}

	keyboard := inline.New(b, inline.NoDeleteAfterClick()).
		Button(constants.PhysicalActivityLow, []byte(constants.PhysicalActivityLowKey), tgbot.onPhysicalActivitySelect).
		Button(constants.PhysicalActivityModerate, []byte(constants.PhysicalActivityModerateKey), tgbot.onPhysicalActivitySelect).
		Button(constants.PhysicalActivityHigh, []byte(constants.PhysicalActivityHighKey), tgbot.onPhysicalActivitySelect)

	_, err = b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      chatID,
		Text:        constants.PhysicalActivityQuestion,
		ReplyMarkup: keyboard,
	})
	if err != nil {
		logger.Get("tgbot", "tgbot.onSexSelect").Error("failed to send message", "error", err.Error())
		return
	}
}

func (tgbot *TGBot) onPhysicalActivitySelect(ctx context.Context, b *bot.Bot, mes models.MaybeInaccessibleMessage, data []byte) {
	chatID := mes.Message.Chat.ID

	err := tgbot.userService.SaveUserDataField(ctx, chatID, constants.PhysicalActivityKey, string(data))
	if err != nil {
		logger.Get("tgbot", "tgbot.onPhysicalActivitySelect").Error("failed to save user data", "error", err.Error())
		return
	}

	keyboard := inline.New(b, inline.NoDeleteAfterClick()).
		Button(constants.ClimateCold, []byte(constants.ClimateColdKey), tgbot.onClimateSelect).
		Button(constants.ClimateTemperate, []byte(constants.ClimateTemperateKey), tgbot.onClimateSelect).
		Button(constants.ClimateWarm, []byte(constants.ClimateWarmKey), tgbot.onClimateSelect).
		Button(constants.ClimateHot, []byte(constants.ClimateHotKey), tgbot.onClimateSelect)

	_, err = b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      chatID,
		Text:        constants.ClimateQuestion,
		ReplyMarkup: keyboard,
	})
	if err != nil {
		logger.Get("tgbot", "tgbot.onPhysicalActivitySelect").Error("failed to send message", "error", err.Error())
		return
	}
}

func (tgbot *TGBot) onClimateSelect(ctx context.Context, b *bot.Bot, mes models.MaybeInaccessibleMessage, data []byte) {
	chatID := mes.Message.Chat.ID

	err := tgbot.userService.SaveUserDataField(ctx, chatID, constants.ClimateKey, string(data))
	if err != nil {
		logger.Get("tgbot", "tgbot.onClimateSelect").Error("failed to save user data", "error", err.Error())
		return
	}

	userData, err := tgbot.userService.CreateUserData(ctx, chatID)
	if err != nil {
		logger.Get("tgbot", "tgbot.onClimateSelect").Error("failed to create user data", "error", err.Error())
		return
	}

	_, err = b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: chatID,
		Text:   fmt.Sprintf(constants.WaterNormResult, userData.WaterGoal),
	})
	if err != nil {
		logger.Get("tgbot", "tgbot.onClimateSelect").Error("failed to send message", "error", err.Error())
		return
	}
}
