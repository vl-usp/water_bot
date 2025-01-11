package model

import (
	"time"

	"github.com/brianvoe/gofakeit/v7"
)

// User is a model of a user.
type User struct {
	ID           int64
	FirstName    string
	LastName     string
	Username     string
	LanguageCode string
	Params       *UserParams
	CreatedAt    time.Time
}

func FakeUser() *User {
	return &User{
		ID:           gofakeit.Int64(),
		FirstName:    gofakeit.FirstName(),
		LastName:     gofakeit.LastName(),
		Username:     gofakeit.Username(),
		LanguageCode: gofakeit.LanguageAbbreviation(),
		CreatedAt:    gofakeit.Date(),
	}
}
