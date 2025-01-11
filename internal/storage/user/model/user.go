package model

import (
	"database/sql"
	"time"
)

// User is a model of a user for the database.
type User struct {
	ID           int64         `db:"id"`
	FirstName    string        `db:"first_name"`
	LastName     string        `db:"last_name"`
	Username     string        `db:"username"`
	LanguageCode string        `db:"language_code"`
	ParamsID     sql.NullInt64 `db:"params_id"`
	CreatedAt    time.Time     `db:"created_at"`
}
