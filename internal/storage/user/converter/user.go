package converter

import (
	"github.com/vl-usp/water_bot/internal/model"
	storageRefsModel "github.com/vl-usp/water_bot/internal/storage/refs/model"
	storageModel "github.com/vl-usp/water_bot/internal/storage/user/model"
)

// ToUserFromStorage converts a user from the storage to the model.
func ToUserFromStorage(user storageModel.User) *model.User {
	if user.ParamsID.Valid {
		return &model.User{
			ID:           user.ID,
			FirstName:    user.FirstName,
			LastName:     user.LastName,
			Username:     user.Username,
			LanguageCode: user.LanguageCode,
			Params:       &model.UserParams{ID: user.ParamsID.Int64},
			CreatedAt:    user.CreatedAt,
		}
	}

	return &model.User{
		ID:           user.ID,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		Username:     user.Username,
		LanguageCode: user.LanguageCode,
		CreatedAt:    user.CreatedAt,
	}
}

// ToFullUserFromStorage converts a user with params from the storage to the model.
func ToFullUserFromStorage(
	user storageModel.User,
	params storageModel.UserParams,
	sex storageRefsModel.Sex,
	physicalActivity storageRefsModel.PhysicalActivity,
	climate storageRefsModel.Climate,
	timezone storageRefsModel.Timezone,
) *model.User {
	return &model.User{
		ID:           user.ID,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		Username:     user.Username,
		LanguageCode: user.LanguageCode,
		Params:       ToUserParamsFromStorage(params, sex, physicalActivity, climate, timezone),
		CreatedAt:    user.CreatedAt,
	}
}
