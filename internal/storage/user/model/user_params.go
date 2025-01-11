package model

import (
	"database/sql"
	"time"
)

// UserParams is a model of user params for the database.
type UserParams struct {
	ID                 int64        `db:"id"`
	SexID              sql.NullByte `db:"sex_id"`
	PhysicalActivityID sql.NullByte `db:"physical_activity_id"`
	ClimateID          sql.NullByte `db:"climate_id"`
	TimezoneID         sql.NullByte `db:"timezone_id"`
	Weight             sql.NullByte `db:"weight"`
	WaterGoal          uint16       `db:"water_goal"`
	CreatedAt          time.Time    `db:"created_at"`
	UpdatedAt          sql.NullTime `db:"updated_at"`
}
