package model

import (
	"time"

	"github.com/brianvoe/gofakeit/v7"
)

// UserParams is a part of user model.
type UserParams struct {
	ID               int64
	Sex              *Sex
	PhysicalActivity *PhysicalActivity
	Climate          *Climate
	Timezone         *Timezone
	Weight           *byte
	WaterGoal        uint16
	CreatedAt        time.Time
	UpdatedAt        *time.Time
}

// FakeUserParams returns a fake user params.
func FakeUserParams() *UserParams {
	weight := gofakeit.Uint8()
	updatedAt := gofakeit.Date()

	return &UserParams{
		ID:        gofakeit.Int64(),
		Weight:    &weight,
		WaterGoal: gofakeit.Uint16(),
		CreatedAt: gofakeit.Date(),
		UpdatedAt: &updatedAt,
	}
}
