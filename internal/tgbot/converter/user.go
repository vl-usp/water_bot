package converter

import (
	"github.com/go-telegram/bot/models"
	"github.com/vl-usp/water_bot/internal/model"
)

// ToUserFromTGUser converts a models.User to a model.User.
func ToUserFromTGUser(user models.User) model.User {
	return model.User{
		ID:           user.ID,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		Username:     user.Username,
		LanguageCode: user.LanguageCode,
	}
}
