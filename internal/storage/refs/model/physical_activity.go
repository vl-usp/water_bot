package model

// PhysicalActivity is a database representation of physical activity
type PhysicalActivity struct {
	ID        byte    `db:"id"`
	Key       string  `db:"key"`
	Name      string  `db:"name"`
	WaterCoef float64 `db:"water_coef"`
}
