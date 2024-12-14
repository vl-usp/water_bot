package converter

import (
	"errors"
	"strconv"

	"github.com/vl-usp/water_bot/internal/constants"
	"github.com/vl-usp/water_bot/internal/model"

	modelRepo "github.com/vl-usp/water_bot/internal/repository/user_data/model"
)

// ToUserDataFromRepo converts a model.UserData to a modelRepo.UserData
func ToUserDataFromRepo(userData *modelRepo.UserData) *model.UserData {
	return &model.UserData{
		ID:                 userData.ID,
		UserID:             userData.UserID,
		SexID:              model.SexID(userData.SexID),
		Weight:             userData.Weight,
		PhysicalActivityID: model.PhysicalActivityID(userData.PhysicalActivityID),
		ClimateID:          model.ClimateID(userData.ClimateID),
		CreatedAt:          userData.CreatedAt.Time,
		UpdatedAt:          userData.UpdatedAt.Time,
	}
}

// ToUserDataFromCache converts a map[string]string to a model.UserData
func ToUserDataFromCache(data map[string]string) (*model.UserData, error) {
	var userData model.UserData
	for key, value := range data {
		switch key {
		case constants.SexKey:
			userData.SexID = sexFromString(value)
		case constants.PhysicalActivityKey:
			userData.PhysicalActivityID = physicalActivityFromString(value)
		case constants.ClimateKey:
			userData.ClimateID = climateFromString(value)
		case constants.WeightKey:
			weight, err := strconv.Atoi(value)
			if err != nil {
				return nil, err
			}

			userData.Weight = weight
		case constants.WaterGoalKey:
			waterGoal, err := strconv.Atoi(value)
			if err != nil {
				return nil, err
			}

			userData.WaterGoal = waterGoal
		default:
			return nil, errors.New("unknown key: " + key)
		}
	}

	return &userData, nil
}

func sexFromString(sex string) model.SexID {
	switch sex {
	case constants.SexMaleKey:
		return model.Male
	case constants.SexFemaleKey:
		return model.Female
	default:
		return 0
	}
}

func physicalActivityFromString(physicalActivity string) model.PhysicalActivityID {
	switch physicalActivity {
	case constants.PhysicalActivityLowKey:
		return model.Low
	case constants.PhysicalActivityModerateKey:
		return model.Moderate
	case constants.PhysicalActivityHighKey:
		return model.High
	default:
		return 0
	}
}

func climateFromString(climate string) model.ClimateID {
	switch climate {
	case constants.ClimateColdKey:
		return model.Cold
	case constants.ClimateTemperateKey:
		return model.Temperate
	case constants.ClimateWarmKey:
		return model.Warm
	case constants.ClimateHotKey:
		return model.Hot
	default:
		return 0
	}
}
