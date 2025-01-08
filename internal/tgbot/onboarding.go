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
func (client *Client) getOnboardingKeyboard(chatID int64) *inline.Keyboard {
	keyboard := inline.New(client.bot)

	switch client.fsm.Current(chatID) {
	case constants.StateOnboardingWaterNorm:
		keyboard.
			Button(constants.WaterNormSelf, []byte(constants.WaterNormSelfKey), client.onWaterNormSelect).
			Button(constants.WaterNormDefault, []byte(constants.WaterNormDefaultKey), client.onWaterNormSelect)
	case constants.StateOnboardingWeight:
		return nil
	case constants.StateOnboardingSex:
		// get sexes from db
		// create keyboard with sexes (2 sexes in a row)
		keyboard.
			Button(constants.SexMale, []byte(constants.SexMaleKey), client.onSexSelect).
			Button(constants.SexFemale, []byte(constants.SexFemaleKey), client.onSexSelect)
	case constants.StateOnboardingPhysicalActivity:
		// get physical activities from db
		// create keyboard with physical activities (5 physical activities in a row)
		keyboard.
			Button(constants.PhysicalActivityLow, []byte(constants.PhysicalActivityLowKey), client.onPhysicalActivitySelect).
			Button(constants.PhysicalActivityModerate, []byte(constants.PhysicalActivityModerateKey), client.onPhysicalActivitySelect).
			Button(constants.PhysicalActivityHigh, []byte(constants.PhysicalActivityHighKey), client.onPhysicalActivitySelect)
	case constants.StateOnboardingClimate:
		// get climates from db
		// create keyboard with climates (5 climates in a row)
		keyboard.
			Button(constants.ClimateCold, []byte(constants.ClimateColdKey), client.onClimateSelect).
			Button(constants.ClimateTemperate, []byte(constants.ClimateTemperateKey), client.onClimateSelect).
			Button(constants.ClimateWarm, []byte(constants.ClimateWarmKey), client.onClimateSelect).
			Button(constants.ClimateHot, []byte(constants.ClimateHotKey), client.onClimateSelect)
	case constants.StateOnboardingTimezone:
		keyboard.
			Button("test", []byte("test"), client.onTimezoneSelect)
		// get timezones from db
		// create keyboard with timezones (5 timezones in a row)

	}

	return keyboard
}

// startOnboarding starts the onboarding process
// It's sends two messages (benifit of water and standard norm)
// then it suggests the user to answer the water norm question
// and transitions to the water norm state.
// Next see onWaterNormSelect function
//
// Flow: startOnboarding -> onWaterNormSelect -> callbackAskWeight -> onboardingWeightHandler (client.inputHandler)
// -> sexSelectHandler -> onSexSelect -> onPhysicalActivitySelect -> onClimateSelect -> onTimezoneSelect
func (client *Client) startOnboarding(ctx context.Context, _ *bot.Bot, update *models.Update) {
	chatID := update.Message.Chat.ID
	client.setOnboardingFSM()
	client.sendMessage(ctx, chatID, constants.BenifitOfWater)
	client.sendMessage(ctx, chatID, constants.StandardNorm)
	client.fsm.Transition(chatID, constants.StateOnboardingWaterNorm, chatID)
	keyboard := client.getOnboardingKeyboard(chatID)
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
		return
	}

	switch string(data) {
	case constants.WaterNormSelfKey:
		client.fsm.Transition(chatID, constants.StateOnboardingWeight, chatID)
	case constants.WaterNormDefaultKey:
		err := client.userService.SaveUserParam(ctx, chatID, constants.WaterGoalKey, constants.WaterGoalDefault)
		if err != nil {
			logger.Get("tgbot", "client.onWaterNormSelect").Error("failed to save user data", "error", err.Error())
			return
		}

		client.fsm.Transition(chatID, constants.StateOnboardingTimezone, chatID)
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
		logger.Get("tgbot", "client.defaultHandler").Error("failed to save weight", "error", err.Error())
		return
	}

	client.sexSelectHandler(ctx, client.bot, update)
}

// sexSelectHandler called from onboardingWeightHandler after user send the weight
func (client *Client) sexSelectHandler(ctx context.Context, _ *bot.Bot, update *models.Update) {
	chatID := update.Message.Chat.ID
	client.fsm.Transition(chatID, constants.StateOnboardingSex, chatID)
	keyboard := client.getOnboardingKeyboard(chatID)
	client.sendMessageWithKeyboard(ctx, chatID, constants.SexQuestion, keyboard)
}

// onSexSelect handles the sex selection.
// After the answer it calls onPhysicalActivitySelect
func (client *Client) onSexSelect(ctx context.Context, _ *bot.Bot, mes models.MaybeInaccessibleMessage, data []byte) {
	chatID := mes.Message.Chat.ID
	if client.fsm.Current(chatID) != constants.StateOnboardingSex {
		logger.Get("tgbot", "client.onWaterNormSelect").Error("unexpected state", "state", client.fsm.Current(chatID))
		return
	}

	err := client.userService.SaveUserParam(ctx, chatID, constants.SexKey, string(data))
	if err != nil {
		logger.Get("tgbot", "client.onSexSelect").Error("failed to save user data", "error", err.Error())
		return
	}

	client.fsm.Transition(chatID, constants.StateOnboardingPhysicalActivity, chatID)
	keyboard := client.getOnboardingKeyboard(chatID)
	client.sendMessageWithKeyboard(ctx, chatID, constants.PhysicalActivityQuestion, keyboard)
}

// onPhysicalActivitySelect handles the physical activity selection.
// After the answer it calls onClimateSelect
func (client *Client) onPhysicalActivitySelect(ctx context.Context, _ *bot.Bot, mes models.MaybeInaccessibleMessage, data []byte) {
	chatID := mes.Message.Chat.ID
	if client.fsm.Current(chatID) != constants.StateOnboardingPhysicalActivity {
		logger.Get("tgbot", "client.onPhysicalActivitySelect").Error("unexpected state", "state", client.fsm.Current(chatID))
		return
	}

	err := client.userService.SaveUserParam(ctx, chatID, constants.PhysicalActivityKey, string(data))
	if err != nil {
		logger.Get("tgbot", "client.onPhysicalActivitySelect").Error("failed to save user data", "error", err.Error())
		return
	}

	client.fsm.Transition(chatID, constants.StateOnboardingClimate, chatID)
	keyboard := client.getOnboardingKeyboard(chatID)
	client.sendMessageWithKeyboard(ctx, chatID, constants.ClimateQuestion, keyboard)
}

// onClimateSelect handles the climate selection.
// After the answer it calls onTimezoneSelect
func (client *Client) onClimateSelect(ctx context.Context, _ *bot.Bot, mes models.MaybeInaccessibleMessage, data []byte) {
	chatID := mes.Message.Chat.ID
	if client.fsm.Current(chatID) != constants.StateOnboardingClimate {
		logger.Get("tgbot", "client.onWaterNormSelect").Error("unexpected state", "state", client.fsm.Current(chatID))
		return
	}

	err := client.userService.SaveUserParam(ctx, chatID, constants.ClimateKey, string(data))
	if err != nil {
		logger.Get("tgbot", "client.onClimateSelect").Error("failed to save user data", "error", err.Error())
		return
	}

	client.fsm.Transition(chatID, constants.StateOnboardingTimezone, chatID)
	keyboard := client.getOnboardingKeyboard(chatID)
	client.sendMessageWithKeyboard(ctx, chatID, constants.TimezoneQuestion, keyboard)
}

// onTimezoneSelect handles the timezone selection.
func (client *Client) onTimezoneSelect(ctx context.Context, _ *bot.Bot, mes models.MaybeInaccessibleMessage, data []byte) {
	chatID := mes.Message.Chat.ID
	if client.fsm.Current(chatID) != constants.StateOnboardingClimate {
		logger.Get("tgbot", "client.onWaterNormSelect").Error("unexpected state", "state", client.fsm.Current(chatID))
		return
	}

	err := client.userService.SaveUserParam(ctx, chatID, constants.TimezoneKey, string(data))
	if err != nil {
		logger.Get("tgbot", "client.onTimezoneSelect").Error("failed to save user data", "error", err.Error())
		return
	}
	client.fsm.Transition(chatID, constants.StateDefault, chatID)

	// save user to database
	err = client.userService.UpdateUserFromCache(ctx, chatID)
	if err != nil {
		logger.Get("tgbot", "client.saveUser").Error("failed to create user data", "error", err.Error())
		return
	}

	user, err := client.userService.GetUser(ctx, chatID)
	if err != nil {
		logger.Get("tgbot", "client.saveUser").Error("failed to create user data", "error", err.Error())
		return
	}
	client.sendMessage(ctx, chatID, fmt.Sprintf(constants.WaterNormResult, user.Params.WaterGoal))
}
