package model

import "time"

type User struct {
	ID           int64
	FirstName    string
	LastName     string
	Username     string
	LanguageCode string
	CreatedAt    time.Time
}
