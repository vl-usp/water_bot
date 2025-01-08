package converter

import (
	"github.com/vl-usp/water_bot/internal/model"
	storageModel "github.com/vl-usp/water_bot/internal/storage/refs/model"
)

// ToTimezoneFromRepo converts a database representation of timezone to model.Timezone
func ToTimezoneFromRepo(r storageModel.Timezone) model.Timezone {
	return model.Timezone{
		ID:        r.ID,
		Name:      r.Name,
		Cities:    r.Cities,
		UTCOffset: r.UTCOffset,
	}
}
