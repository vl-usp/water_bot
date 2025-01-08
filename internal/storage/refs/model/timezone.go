package model

// Timezone is a database representation of timezone
type Timezone struct {
	ID        byte   `db:"id"`
	Name      string `db:"name"`
	Cities    string `db:"cities"`
	UTCOffset int    `db:"utc_offset"`
}
