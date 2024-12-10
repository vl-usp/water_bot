package model

import "time"

// User is a model of a user.
type User struct {
	ID           int64
	FirstName    string
	LastName     string
	Username     string
	LanguageCode string
	CreatedAt    time.Time
}
