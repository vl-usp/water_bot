package model

import (
	"github.com/brianvoe/gofakeit/v7"
)

// Climate is a model of climate
type Climate struct {
	ID        byte
	Key       string
	Name      string
	WaterCoef float64
}

func FakeClimate() *Climate {
	return &Climate{
		ID:        gofakeit.Uint8(),
		Key:       gofakeit.Word(),
		Name:      gofakeit.Word(),
		WaterCoef: gofakeit.Float64(),
	}
}
