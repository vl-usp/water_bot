package converter

import (
	"errors"
	"strconv"

	"github.com/vl-usp/water_bot/internal/constants"
	"github.com/vl-usp/water_bot/internal/model"
)

// ToUserParamsFromCache converts a map[string]string to a model.UserParams
func ToUserParamsFromCache(data map[string]string) (*model.UserParams, error) {
	var userData model.UserParams
	for key, value := range data {
		switch key {
		case constants.SexKey:
			id, err := strconv.Atoi(value)
			if err != nil {
				return nil, err
			}
			userData.Sex = &model.Sex{ID: byte(id)}

		case constants.PhysicalActivityKey:
			id, err := strconv.Atoi(value)
			if err != nil {
				return nil, err
			}
			userData.PhysicalActivity = &model.PhysicalActivity{ID: byte(id)}

		case constants.ClimateKey:
			id, err := strconv.Atoi(value)
			if err != nil {
				return nil, err
			}
			userData.Climate = &model.Climate{ID: byte(id)}

		case constants.TimezoneKey:
			id, err := strconv.Atoi(value)
			if err != nil {
				return nil, err
			}
			userData.Timezone = &model.Timezone{ID: byte(id)}

		case constants.WeightKey:
			weight, err := strconv.Atoi(value)
			if err != nil {
				return nil, err
			}

			weightByte := byte(weight)
			userData.Weight = &weightByte

		case constants.WaterGoalKey:
			waterGoal, err := strconv.Atoi(value)
			if err != nil {
				return nil, err
			}

			userData.WaterGoal = uint16(waterGoal)

		default:
			return nil, errors.New("unknown key: " + key)
		}
	}

	return &userData, nil
}
