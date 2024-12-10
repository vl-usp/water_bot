package converter

import (
	"github.com/vl-usp/water_bot/internal/model"
	modelRepo "github.com/vl-usp/water_bot/internal/repository/user/model"
)

func ToUserFromRepo(user *modelRepo.User) *model.User {
	return &model.User{
		ID:           user.ID,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		LanguageCode: user.LanguageCode,
		CreatedAt:    user.CreatedAt.Time,
	}
}
