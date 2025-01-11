package model

import "database/sql"

// PhysicalActivity is a database representation of physical activity
type PhysicalActivity struct {
	ID        sql.NullByte    `db:"id"`
	Key       sql.NullString  `db:"key"`
	Name      sql.NullString  `db:"name"`
	WaterCoef sql.NullFloat64 `db:"water_coef"`
}
