package model

import "database/sql"

// User is a model of a user for the database.
type User struct {
	ID           int64        `db:"id"`
	FirstName    string       `db:"first_name"`
	LastName     string       `db:"last_name"`
	Username     string       `db:"username"`
	LanguageCode string       `db:"language_code"`
	ParamsID     int64        `db:"params_id"`
	CreatedAt    sql.NullTime `db:"created_at"`
}
