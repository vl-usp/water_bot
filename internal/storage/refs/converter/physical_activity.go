package converter

import (
	"github.com/vl-usp/water_bot/internal/model"
	storageModel "github.com/vl-usp/water_bot/internal/storage/refs/model"
)

// ToPhysicalActivityFromRepo converts a database representation of physical activity to model.PhysicalActivity
func ToPhysicalActivityFromRepo(r storageModel.PhysicalActivity) model.PhysicalActivity {
	return model.PhysicalActivity{
		ID:        r.ID,
		Key:       r.Key,
		Name:      r.Name,
		WaterCoef: r.WaterCoef,
	}
}
