package converter

import (
	"github.com/vl-usp/water_bot/internal/model"
	storageModel "github.com/vl-usp/water_bot/internal/storage/refs/model"
)

// ToClimateFromRepo converts a database representation of climate to model.Climate
func ToClimateFromRepo(r storageModel.Climate) model.Climate {
	return model.Climate{
		ID:        r.ID,
		Key:       r.Key,
		Name:      r.Name,
		WaterCoef: r.WaterCoef,
	}
}
