package converter

import (
	"github.com/vl-usp/water_bot/internal/model"
	storageRefsModel "github.com/vl-usp/water_bot/internal/storage/refs/model"
	storageModel "github.com/vl-usp/water_bot/internal/storage/user/model"
)

// ToUserFromRepo converts a user from the repository to the model.
func ToUserFromRepo(
	user storageModel.User,
	params storageModel.UserParams,
	sex storageRefsModel.Sex,
	physicalActivity storageRefsModel.PhysicalActivity,
	climate storageRefsModel.Climate,
	timezone storageRefsModel.Timezone,
) model.User {
	return model.User{
		ID:           user.ID,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		Username:     user.Username,
		LanguageCode: user.LanguageCode,
		Params:       ToUserParamsFromRepo(params, sex, physicalActivity, climate, timezone),
		CreatedAt:    user.CreatedAt.Time,
	}
}
