package model

// Climate is a database representation of climate
type Climate struct {
	ID        byte    `db:"id"`
	Key       string  `db:"key"`
	Name      string  `db:"name"`
	WaterCoef float64 `db:"water_coef"`
}
