package model

import "database/sql"

// UserParams is a model of user params for the database.
type UserParams struct {
	ID                 int64        `db:"id"`
	SexID              byte         `db:"sex_id"`
	PhysicalActivityID byte         `db:"physical_activity_id"`
	ClimateID          byte         `db:"climate_id"`
	TimezoneID         byte         `db:"timezone_id"`
	Weight             byte         `db:"weight"`
	WaterGoal          int          `db:"water_goal"`
	CreatedAt          sql.NullTime `db:"created_at"`
	UpdatedAt          sql.NullTime `db:"updated_at"`
}
