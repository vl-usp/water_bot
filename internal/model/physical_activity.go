package model

import (
	"github.com/brianvoe/gofakeit/v7"
)

// PhysicalActivity is a model of physical activity
type PhysicalActivity struct {
	ID        byte
	Key       string
	Name      string
	WaterCoef float64
}

// FakePhysicalActivity returns a fake physical activity
func FakePhysicalActivity() *PhysicalActivity {
	return &PhysicalActivity{
		ID:        gofakeit.Uint8(),
		Key:       gofakeit.Word(),
		Name:      gofakeit.Word(),
		WaterCoef: gofakeit.Float64(),
	}
}
