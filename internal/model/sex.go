package model

import (
	"github.com/brianvoe/gofakeit/v7"
)

// Sex is a model of sex
type Sex struct {
	ID        byte
	Key       string
	Name      string
	WaterCoef float64
}

func FakeSex() *Sex {
	return &Sex{
		ID:        gofakeit.Uint8(),
		Key:       gofakeit.Word(),
		Name:      gofakeit.Word(),
		WaterCoef: gofakeit.Float64(),
	}
}
