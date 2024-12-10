package errors

import (
	"errors"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
)

func IsUniqueViolation(err error) bool {
	var e *pgconn.PgError
	return errors.As(err, &e) && e.Code == pgerrcode.UniqueViolation
}
