package db

import (
	"context"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

// Handler is a function that is executed in a transaction
type Handler func(ctx context.Context) error

// Client represents a database client
type Client interface {
	DB() DB
	Close() error
}

// TxManager executes the specified user handler in a transaction
type TxManager interface {
	ReadCommitted(ctx context.Context, f Handler) error
}

// Query is a wrapper around a query, storing the name of the query and the query itself
// The name of the query is used for logging and may be used elsewhere, for example, for tracing
type Query struct {
	Name     string
	QueryRaw string
}

// Transactor is an interface for working with transactions
type Transactor interface {
	BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error)
}

// SQLExecer is a combination of NamedExecer and QueryExecer
type SQLExecer interface {
	NamedExecer
	QueryExecer
}

// NamedExecer is an interface for working with named queries with tags in structs
type NamedExecer interface {
	ScanOneContext(ctx context.Context, dest interface{}, q Query, args ...interface{}) error
	ScanAllContext(ctx context.Context, dest interface{}, q Query, args ...interface{}) error
}

// QueryExecer is an interface for working with ordinary queries
type QueryExecer interface {
	ExecContext(ctx context.Context, q Query, args ...interface{}) (pgconn.CommandTag, error)
	QueryContext(ctx context.Context, q Query, args ...interface{}) (pgx.Rows, error)
	QueryRowContext(ctx context.Context, q Query, args ...interface{}) pgx.Row
}

// Pinger is an interface for checking the connection to the database
type Pinger interface {
	Ping(ctx context.Context) error
}

// DB is an interface for working with a database
type DB interface {
	SQLExecer
	Transactor
	Pinger
	Close()
}
