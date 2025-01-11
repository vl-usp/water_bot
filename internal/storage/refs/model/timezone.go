package model

import "database/sql"

// Timezone is a database representation of timezone
type Timezone struct {
	ID        sql.NullByte   `db:"id"`
	Name      sql.NullString `db:"name"`
	Cities    sql.NullString `db:"cities"`
	UTCOffset sql.NullInt16  `db:"utc_offset"`
}
