package tgbot

import (
	"context"
	"fmt"
	"strconv"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/go-telegram/fsm"
	"github.com/go-telegram/ui/keyboard/inline"
	"github.com/vl-usp/water_bot/internal/constants"
	"github.com/vl-usp/water_bot/pkg/logger"
)

func (client *Client) setOnboardingFSM() {
	client.fsm = fsm.New(
		constants.StateDefault,
		map[fsm.StateID]fsm.Callback{
			constants.StateOnboardingWeight: client.callbackAskWeight,
		},
	)
}

// getOnboardingKeyboard returns the keyboard for the onboarding by the current state
func (client *Client) getOnboardingKeyboard(ctx context.Context, chatID int64) (*inline.Keyboard, error) {
	keyboard := inline.New(client.bot)

	switch client.fsm.Current(chatID) {
	case constants.StateOnboardingWaterNorm:
		keyboard.
			Button(constants.WaterNormSelf, []byte(constants.WaterNormSelfKey), client.onWaterNormSelect).
			Button(constants.WaterNormDefault, []byte(constants.WaterNormDefaultKey), client.onWaterNormSelect)
	case constants.StateOnboardingWeight:
		return nil, nil
	case constants.StateOnboardingSex:
		list, err := client.refProvider.SexList(ctx)
		if err != nil {
			return nil, err
		}

		for _, item := range list {
			keyboard.Button(item.Name, []byte{item.ID}, client.onSexSelect)
		}
	case constants.StateOnboardingPhysicalActivity:
		list, err := client.refProvider.PhysicalActivityList(ctx)
		if err != nil {
			return nil, err
		}

		for _, item := range list {
			keyboard.Button(item.Name, []byte{item.ID}, client.onPhysicalActivitySelect)
		}
	case constants.StateOnboardingClimate:
		list, err := client.refProvider.ClimateList(ctx)
		if err != nil {
			return nil, err
		}

		for _, item := range list {
			keyboard.Button(item.Name, []byte{item.ID}, client.onClimateSelect)
		}
	case constants.StateOnboardingTimezone:
		// TODO: передалеть, чтобы можно было выводить текст таблицу в сообщении
		// Придумать универсальный вариант
		list, err := client.refProvider.TimezoneList(ctx)
		if err != nil {
			return nil, err
		}

		for i, item := range list {
			keyboard.Button(item.Name, []byte{item.ID}, client.onTimezoneSelect)
			if (i+1)%5 == 0 {
				keyboard.Row()
			}
		}
	}

	return keyboard, nil
}

// startOnboarding starts the onboarding process
// It's sends two messages (benifit of water and standard norm)
// then it suggests the user to answer the water norm question
// and transitions to the water norm state.
//
// Flow: startOnboarding -> onWaterNormSelect -> callbackAskWeight -> onboardingWeightHandler (client.inputHandler)
// -> sexSelectHandler -> onSexSelect -> onPhysicalActivitySelect -> onClimateSelect -> onTimezoneSelect
func (client *Client) startOnboarding(ctx context.Context, _ *bot.Bot, update *models.Update) {
	chatID := update.Message.Chat.ID
	client.setOnboardingFSM()
	client.sendMessage(ctx, chatID, constants.BenifitOfWater)
	client.sendMessage(ctx, chatID, constants.StandardNorm)
	client.fsm.Transition(chatID, constants.StateOnboardingWaterNorm, chatID)
	keyboard, err := client.getOnboardingKeyboard(ctx, chatID)
	if err != nil {
		logger.Get("tgbot", "client.startOnboarding").Error("failed to get keyboard", "error", err.Error())
		client.sendErrorMessage(ctx, update.Message.Chat.ID, constants.DefaultErrorMessage)
		return
	}
	client.sendMessageWithKeyboard(ctx, chatID, constants.WaterNormQuestion, keyboard)
}

// onWaterNormSelect handles the water norm selection.
// User have two options: self or default.
// If user selects self, it transitions to the weight state.
// If user selects default, it saves the water goal to db.
func (client *Client) onWaterNormSelect(ctx context.Context, _ *bot.Bot, mes models.MaybeInaccessibleMessage, data []byte) {
	chatID := mes.Message.Chat.ID
	if client.fsm.Current(chatID) != constants.StateOnboardingWaterNorm {
		logger.Get("tgbot", "client.onWaterNormSelect").Error("unexpected state", "state", client.fsm.Current(chatID))
		client.sendErrorMessage(ctx, chatID, constants.DefaultErrorMessage)
		return
	}

	switch string(data) {
	case constants.WaterNormSelfKey:
		client.fsm.Transition(chatID, constants.StateOnboardingWeight, chatID)
	case constants.WaterNormDefaultKey:
		err := client.userService.SaveUserParam(ctx, chatID, constants.WaterGoalKey, constants.WaterGoalDefault)
		if err != nil {
			logger.Get("tgbot", "client.onWaterNormSelect").Error("failed to save user data", "error", err.Error())
			client.sendErrorMessage(ctx, chatID, constants.DefaultErrorMessage)
			return
		}

		client.fsm.Transition(chatID, constants.StateOnboardingTimezone, chatID)
		keyboard, err := client.getOnboardingKeyboard(ctx, chatID)
		if err != nil {
			logger.Get("tgbot", "client.onWaterNormSelect").Error("failed to get keyboard", "error", err.Error())
			client.sendErrorMessage(ctx, chatID, constants.DefaultErrorMessage)
			return
		}
		client.sendMessageWithKeyboard(ctx, chatID, constants.TimezoneKey, keyboard)
	}
}

// askWeightHandler sends the weight question and wait for the text answer, that will be handled by onboardingWeightHandler
func (client *Client) callbackAskWeight(_ *fsm.FSM, args ...any) {
	chatID := args[0]
	client.sendMessage(context.Background(), chatID.(int64), constants.WeightQuestion)
}

// onboardingWeightHandler handles the weight selection.
// This function called from client.inputHandler, because user send the weight as text
func (client *Client) onboardingWeightHandler(ctx context.Context, _ *bot.Bot, update *models.Update) {
	chatID := update.Message.Chat.ID
	if client.fsm.Current(chatID) != constants.StateOnboardingWeight {
		logger.Get("tgbot", "client.onboardingWeightHandler").Error("unexpected state", "state", client.fsm.Current(chatID))
		client.sendErrorMessage(ctx, chatID, constants.DefaultErrorMessage)
		return
	}

	weight, errWeight := strconv.Atoi(update.Message.Text)
	if errWeight != nil {
		client.sendMessage(ctx, chatID, constants.WeightErrorMessage)
		return
	}
	if weight < 30 || weight > 300 {
		client.sendMessage(ctx, chatID, constants.WeightNotValidMessage)
		return
	}

	// Store the weight to cache
	err := client.userService.SaveUserParam(ctx, chatID, constants.WeightKey, weight)
	if err != nil {
		logger.Get("tgbot", "client.defaultHandler").Error("failed to save user param", "error", err.Error())
		client.sendErrorMessage(ctx, chatID, constants.DefaultErrorMessage)
		return
	}

	client.sexSelectHandler(ctx, client.bot, update)
}

// sexSelectHandler called from onboardingWeightHandler after user send the weight
func (client *Client) sexSelectHandler(ctx context.Context, _ *bot.Bot, update *models.Update) {
	chatID := update.Message.Chat.ID
	client.fsm.Transition(chatID, constants.StateOnboardingSex, chatID)
	keyboard, err := client.getOnboardingKeyboard(ctx, chatID)
	if err != nil {
		logger.Get("tgbot", "client.sexSelectHandler").Error("failed to get keyboard", "error", err.Error())
		client.sendErrorMessage(ctx, chatID, constants.DefaultErrorMessage)
		return
	}
	client.sendMessageWithKeyboard(ctx, chatID, constants.SexQuestion, keyboard)
}

// onSexSelect handles the sex selection.
// After the answer it calls onPhysicalActivitySelect
func (client *Client) onSexSelect(ctx context.Context, _ *bot.Bot, mes models.MaybeInaccessibleMessage, data []byte) {
	chatID := mes.Message.Chat.ID
	if client.fsm.Current(chatID) != constants.StateOnboardingSex {
		logger.Get("tgbot", "client.onSexSelect").Error("unexpected state", "state", client.fsm.Current(chatID))
		client.sendErrorMessage(ctx, chatID, constants.DefaultErrorMessage)
		return
	}

	sexID := data[0]
	err := client.userService.SaveUserParam(ctx, chatID, constants.SexKey, sexID)
	if err != nil {
		logger.Get("tgbot", "client.onSexSelect").Error("failed to save user param", "error", err.Error())
		client.sendErrorMessage(ctx, chatID, constants.DefaultErrorMessage)
		return
	}

	client.fsm.Transition(chatID, constants.StateOnboardingPhysicalActivity, chatID)
	keyboard, err := client.getOnboardingKeyboard(ctx, chatID)
	if err != nil {
		logger.Get("tgbot", "client.onSexSelect").Error("failed to get keyboard", "error", err.Error())
		client.sendErrorMessage(ctx, chatID, constants.DefaultErrorMessage)
		return
	}
	client.sendMessageWithKeyboard(ctx, chatID, constants.PhysicalActivityQuestion, keyboard)
}

// onPhysicalActivitySelect handles the physical activity selection.
// After the answer it calls onClimateSelect
func (client *Client) onPhysicalActivitySelect(ctx context.Context, _ *bot.Bot, mes models.MaybeInaccessibleMessage, data []byte) {
	chatID := mes.Message.Chat.ID
	if client.fsm.Current(chatID) != constants.StateOnboardingPhysicalActivity {
		logger.Get("tgbot", "client.onPhysicalActivitySelect").Error("unexpected state", "state", client.fsm.Current(chatID))
		client.sendErrorMessage(ctx, chatID, constants.DefaultErrorMessage)
		return
	}

	physicalActivityID := data[0]
	err := client.userService.SaveUserParam(ctx, chatID, constants.PhysicalActivityKey, physicalActivityID)
	if err != nil {
		logger.Get("tgbot", "client.onPhysicalActivitySelect").Error("failed to save user param", "error", err.Error())
		client.sendErrorMessage(ctx, chatID, constants.DefaultErrorMessage)
		return
	}

	client.fsm.Transition(chatID, constants.StateOnboardingClimate, chatID)
	keyboard, err := client.getOnboardingKeyboard(ctx, chatID)
	if err != nil {
		logger.Get("tgbot", "client.onPhysicalActivitySelect").Error("failed to get keyboard", "error", err.Error())
		client.sendErrorMessage(ctx, chatID, constants.DefaultErrorMessage)
		return
	}
	client.sendMessageWithKeyboard(ctx, chatID, constants.ClimateQuestion, keyboard)
}

// onClimateSelect handles the climate selection.
// After the answer it calls onTimezoneSelect
func (client *Client) onClimateSelect(ctx context.Context, _ *bot.Bot, mes models.MaybeInaccessibleMessage, data []byte) {
	chatID := mes.Message.Chat.ID
	if client.fsm.Current(chatID) != constants.StateOnboardingClimate {
		logger.Get("tgbot", "client.onClimateSelect").Error("unexpected state", "state", client.fsm.Current(chatID))
		client.sendErrorMessage(ctx, chatID, constants.DefaultErrorMessage)
		return
	}

	climateID := data[0]
	err := client.userService.SaveUserParam(ctx, chatID, constants.ClimateKey, climateID)
	if err != nil {
		logger.Get("tgbot", "client.onClimateSelect").Error("failed to save user param", "error", err.Error())
		client.sendErrorMessage(ctx, chatID, constants.DefaultErrorMessage)
		return
	}

	client.fsm.Transition(chatID, constants.StateOnboardingTimezone, chatID)
	keyboard, err := client.getOnboardingKeyboard(ctx, chatID)
	if err != nil {
		logger.Get("tgbot", "client.onClimateSelect").Error("failed to get keyboard", "error", err.Error())
		client.sendErrorMessage(ctx, chatID, constants.DefaultErrorMessage)
		return
	}
	client.sendMessageWithKeyboard(ctx, chatID, constants.TimezoneQuestion, keyboard)
}

// onTimezoneSelect handles the timezone selection.
func (client *Client) onTimezoneSelect(ctx context.Context, _ *bot.Bot, mes models.MaybeInaccessibleMessage, data []byte) {
	chatID := mes.Message.Chat.ID
	if client.fsm.Current(chatID) != constants.StateOnboardingTimezone {
		logger.Get("tgbot", "client.onTimezoneSelect").Error("unexpected state", "state", client.fsm.Current(chatID))
		client.sendErrorMessage(ctx, chatID, constants.DefaultErrorMessage)
		return
	}

	timezoneID := data[0]
	err := client.userService.SaveUserParam(ctx, chatID, constants.TimezoneKey, timezoneID)
	if err != nil {
		logger.Get("tgbot", "client.onTimezoneSelect").Error("failed to save user param", "error", err.Error())
		client.sendErrorMessage(ctx, chatID, constants.DefaultErrorMessage)
		return
	}

	// save user to database
	err = client.userService.UpdateUserFromCache(ctx, chatID)
	if err != nil {
		logger.Get("tgbot", "client.onTimezoneSelect").Error("failed to update user from cache", "error", err.Error())
		client.sendErrorMessage(ctx, chatID, constants.DefaultErrorMessage)
		return
	}

	user, err := client.userService.GetFullUser(ctx, chatID)
	if err != nil {
		logger.Get("tgbot", "client.onTimezoneSelect").Error("failed to get user", "error", err.Error())
		client.sendErrorMessage(ctx, chatID, constants.DefaultErrorMessage)
		return
	}
	client.fsm.Transition(chatID, constants.StateDefault, chatID)
	client.sendMessage(ctx, chatID, fmt.Sprintf(constants.WaterNormResult, user.Params.WaterGoal))
}
