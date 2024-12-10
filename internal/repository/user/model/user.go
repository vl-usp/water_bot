package model

import "database/sql"

type User struct {
	ID           int64
	FirstName    string
	LastName     string
	Username     string
	LanguageCode string
	CreatedAt    sql.NullTime
}
