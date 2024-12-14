package model

import "database/sql"

// UserData is a model of user data for the database.
type UserData struct {
	ID                 int64
	UserID             int64
	SexID              int
	PhysicalActivityID int
	ClimateID          int
	Weight             int
	WaterGoal          int
	CreatedAt          sql.NullTime
	UpdatedAt          sql.NullTime
}

// UserDataCache represents a cache for user data.
type UserDataCache struct {
	Sex              string `redis:"sex"`
	PhysicalActivity string `redis:"physical_activity"`
	Climate          string `redis:"climate"`
	Weight           int    `redis:"weight"`
	WaterGoal        int    `redis:"water_goal"`
}
