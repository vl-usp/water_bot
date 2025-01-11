package model

import "database/sql"

// Climate is a database representation of climate
type Climate struct {
	ID        sql.NullByte    `db:"id"`
	Key       sql.NullString  `db:"key"`
	Name      sql.NullString  `db:"name"`
	WaterCoef sql.NullFloat64 `db:"water_coef"`
}
