package converter

import (
	"github.com/vl-usp/water_bot/internal/model"
	storageModel "github.com/vl-usp/water_bot/internal/storage/refs/model"
)

// ToPhysicalActivityFromStorage converts a database representation of physical activity to model.PhysicalActivity
func ToPhysicalActivityFromStorage(from storageModel.PhysicalActivity) *model.PhysicalActivity {
	if !from.ID.Valid {
		return nil
	}

	return &model.PhysicalActivity{
		ID:        from.ID.Byte,
		Key:       from.Key.String,
		Name:      from.Name.String,
		WaterCoef: from.WaterCoef.Float64,
	}
}
