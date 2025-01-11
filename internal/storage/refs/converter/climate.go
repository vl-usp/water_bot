package converter

import (
	"github.com/vl-usp/water_bot/internal/model"
	storageModel "github.com/vl-usp/water_bot/internal/storage/refs/model"
)

// ToClimateFromStorage converts a database representation of climate to model.Climate
func ToClimateFromStorage(from storageModel.Climate) *model.Climate {
	if !from.ID.Valid {
		return nil
	}

	return &model.Climate{
		ID:        from.ID.Byte,
		Key:       from.Key.String,
		Name:      from.Name.String,
		WaterCoef: from.WaterCoef.Float64,
	}
}
