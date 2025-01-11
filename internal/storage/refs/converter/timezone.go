package converter

import (
	"github.com/vl-usp/water_bot/internal/model"
	storageModel "github.com/vl-usp/water_bot/internal/storage/refs/model"
)

// ToTimezoneFromStorage converts a database representation of timezone to model.Timezone
func ToTimezoneFromStorage(from storageModel.Timezone) *model.Timezone {
	if !from.ID.Valid {
		return nil
	}

	return &model.Timezone{
		ID:        from.ID.Byte,
		Name:      from.Name.String,
		Cities:    from.Cities.String,
		UTCOffset: from.UTCOffset.Int16,
	}
}
