package model

import "database/sql"

// Sex is a database representation of sex
type Sex struct {
	ID        sql.NullByte    `db:"id"`
	Key       sql.NullString  `db:"key"`
	Name      sql.NullString  `db:"name"`
	WaterCoef sql.NullFloat64 `db:"water_coef"`
}
