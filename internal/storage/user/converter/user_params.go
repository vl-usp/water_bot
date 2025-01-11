package converter

import (
	"github.com/vl-usp/water_bot/internal/model"
	refsConverter "github.com/vl-usp/water_bot/internal/storage/refs/converter"
	storageRefsModel "github.com/vl-usp/water_bot/internal/storage/refs/model"
	storageModel "github.com/vl-usp/water_bot/internal/storage/user/model"
)

// ToUserParamsFromStorage converts a database representation of user params to a model.UserParams.
func ToUserParamsFromStorage(
	params storageModel.UserParams,
	sex storageRefsModel.Sex,
	physicalActivity storageRefsModel.PhysicalActivity,
	climate storageRefsModel.Climate,
	timezone storageRefsModel.Timezone,
) *model.UserParams {
	res := &model.UserParams{
		ID:               params.ID,
		Sex:              refsConverter.ToSexFromStorage(sex),
		PhysicalActivity: refsConverter.ToPhysicalActivityFromStorage(physicalActivity),
		Climate:          refsConverter.ToClimateFromStorage(climate),
		Timezone:         refsConverter.ToTimezoneFromStorage(timezone),
		WaterGoal:        params.WaterGoal,
		CreatedAt:        params.CreatedAt,
	}

	if params.Weight.Valid {
		res.Weight = &params.Weight.Byte
	}

	if params.UpdatedAt.Valid {
		res.UpdatedAt = &params.UpdatedAt.Time
	}

	return res
}
