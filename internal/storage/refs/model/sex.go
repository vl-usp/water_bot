package model

// Sex is a database representation of sex
type Sex struct {
	ID        byte    `db:"id"`
	Key       string  `db:"key"`
	Name      string  `db:"name"`
	WaterCoef float64 `db:"water_coef"`
}
