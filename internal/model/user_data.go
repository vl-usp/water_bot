package model

import "time"

// SexID is a value type for sex.
type SexID int

// PhysicalActivityID is a value type for physical activity.
type PhysicalActivityID int

// ClimateID is a value type for climate.
type ClimateID int

// SexID enum
const (
	Male SexID = 1 + iota
	Female
)

// PhysicalActivityID enum
const (
	Low PhysicalActivityID = 1 + iota
	Moderate
	High
)

// ClimateID enum
const (
	Cold ClimateID = 1 + iota
	Temperate
	Warm
	Hot
)

// UserData is a model of user data.
type UserData struct {
	ID                 int64
	UserID             int64
	SexID              SexID
	PhysicalActivityID PhysicalActivityID
	ClimateID          ClimateID
	Weight             int
	WaterGoal          int
	CreatedAt          time.Time
	UpdatedAt          time.Time
}
