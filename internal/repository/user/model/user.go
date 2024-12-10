package model

import "database/sql"

// User is a model of a user for the database.
type User struct {
	ID           int64
	FirstName    string
	LastName     string
	Username     string
	LanguageCode string
	CreatedAt    sql.NullTime
}
