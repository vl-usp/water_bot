package model

import "time"

// User is a model of a user.
type User struct {
	ID           int64
	FirstName    string
	LastName     string
	Username     string
	LanguageCode string
	Params       UserParams
	CreatedAt    time.Time
}

// UserParams is a part of user model.
type UserParams struct {
	ID               int64
	Sex              Sex
	PhysicalActivity PhysicalActivity
	Climate          Climate
	Timezone         Timezone
	Weight           byte
	WaterGoal        int
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
