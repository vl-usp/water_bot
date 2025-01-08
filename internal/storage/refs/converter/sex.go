package converter

import (
	"github.com/vl-usp/water_bot/internal/model"
	storageModel "github.com/vl-usp/water_bot/internal/storage/refs/model"
)

// ToSexFromRepo converts a database representation of sex to model.Sex
func ToSexFromRepo(r storageModel.Sex) model.Sex {
	return model.Sex{
		ID:        r.ID,
		Key:       r.Key,
		Name:      r.Name,
		WaterCoef: r.WaterCoef,
	}
}
