package converter

import (
	"github.com/vl-usp/water_bot/internal/model"
	storageModel "github.com/vl-usp/water_bot/internal/storage/refs/model"
)

// ToSexFromStorage converts a database representation of sex to model.Sex
func ToSexFromStorage(from storageModel.Sex) *model.Sex {
	if !from.ID.Valid {
		return nil
	}

	return &model.Sex{
		ID:        from.ID.Byte,
		Key:       from.Key.String,
		Name:      from.Name.String,
		WaterCoef: from.WaterCoef.Float64,
	}
}
