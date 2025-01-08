package converter

import (
	"github.com/vl-usp/water_bot/internal/model"
	refsConverter "github.com/vl-usp/water_bot/internal/storage/refs/converter"
	storageRefsModel "github.com/vl-usp/water_bot/internal/storage/refs/model"
	storageModel "github.com/vl-usp/water_bot/internal/storage/user/model"
)

// ToUserParamsFromRepo converts a database representation of user params to a model.UserParams.
func ToUserParamsFromRepo(
	params storageModel.UserParams,
	sex storageRefsModel.Sex,
	physicalActivity storageRefsModel.PhysicalActivity,
	climate storageRefsModel.Climate,
	timezone storageRefsModel.Timezone,
) model.UserParams {
	return model.UserParams{
		ID:               params.ID,
		Sex:              refsConverter.ToSexFromRepo(sex),
		PhysicalActivity: refsConverter.ToPhysicalActivityFromRepo(physicalActivity),
		Climate:          refsConverter.ToClimateFromRepo(climate),
		Timezone:         refsConverter.ToTimezoneFromRepo(timezone),
		Weight:           params.Weight,
		WaterGoal:        params.WaterGoal,
		CreatedAt:        params.CreatedAt.Time,
		UpdatedAt:        params.UpdatedAt.Time,
	}
}
