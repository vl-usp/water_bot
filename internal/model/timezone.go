package model

import (
	"github.com/brianvoe/gofakeit/v7"
)

// Timezone is a model that represents timezone
type Timezone struct {
	ID        byte
	Name      string
	Cities    string
	UTCOffset int16
}

// FakeTimezone returns a fake timezone
func FakeTimezone() *Timezone {
	return &Timezone{
		ID:        gofakeit.Uint8(),
		Name:      gofakeit.Word(),
		Cities:    gofakeit.Word(),
		UTCOffset: gofakeit.Int16(),
	}
}
