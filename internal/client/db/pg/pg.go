package pg

import (
	"context"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/vl-usp/water_bot/internal/client/db"
	"github.com/vl-usp/water_bot/internal/client/db/prettier"
	"github.com/vl-usp/water_bot/pkg/logger"
)

type key string

// TxKey is the transaction key.
const (
	TxKey key = "tx"
)

type pg struct {
	dbc *pgxpool.Pool
}

// NewDB creates a new DB client.
func NewDB(dbc *pgxpool.Pool) db.DB {
	return &pg{
		dbc: dbc,
	}
}

// ScanOneContext scans one row from the database.
func (p *pg) ScanOneContext(ctx context.Context, dest interface{}, q db.Query, args ...interface{}) error {
	logQuery(ctx, q, args...)

	row, err := p.QueryContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return pgxscan.ScanOne(dest, row)
}

// ScanAllContext scans all rows from the database.
func (p *pg) ScanAllContext(ctx context.Context, dest interface{}, q db.Query, args ...interface{}) error {
	logQuery(ctx, q, args...)

	rows, err := p.QueryContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return pgxscan.ScanAll(dest, rows)
}

// ExecContext executes a query.
func (p *pg) ExecContext(ctx context.Context, q db.Query, args ...interface{}) (pgconn.CommandTag, error) {
	logQuery(ctx, q, args...)

	tx, ok := ctx.Value(TxKey).(pgx.Tx)
	if ok {
		return tx.Exec(ctx, q.QueryRaw, args...)
	}

	return p.dbc.Exec(ctx, q.QueryRaw, args...)
}

// QueryContext executes a query with a context.
func (p *pg) QueryContext(ctx context.Context, q db.Query, args ...interface{}) (pgx.Rows, error) {
	logQuery(ctx, q, args...)

	tx, ok := ctx.Value(TxKey).(pgx.Tx)
	if ok {
		return tx.Query(ctx, q.QueryRaw, args...)
	}

	return p.dbc.Query(ctx, q.QueryRaw, args...)
}

// QueryRowContext executes a query with a context.
func (p *pg) QueryRowContext(ctx context.Context, q db.Query, args ...interface{}) pgx.Row {
	logQuery(ctx, q, args...)

	tx, ok := ctx.Value(TxKey).(pgx.Tx)
	if ok {
		return tx.QueryRow(ctx, q.QueryRaw, args...)
	}

	return p.dbc.QueryRow(ctx, q.QueryRaw, args...)
}

// BeginTx starts a transaction.
func (p *pg) BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error) {
	return p.dbc.BeginTx(ctx, txOptions)
}

// Ping pings the database.
func (p *pg) Ping(ctx context.Context) error {
	return p.dbc.Ping(ctx)
}

// Close closes the database.
func (p *pg) Close() {
	p.dbc.Close()
}

// MakeContextTx makes a context with a transaction.
func MakeContextTx(ctx context.Context, tx pgx.Tx) context.Context {
	return context.WithValue(ctx, TxKey, tx)
}

func logQuery(ctx context.Context, q db.Query, args ...interface{}) {
	prettyQuery := prettier.Pretty(q.QueryRaw, prettier.PlaceholderDollar, args...)

	logger.Get("pg", "logQuery").Info(
		"pg query log",
		"ctx", ctx,
		"sql", q.Name,
		"query", prettyQuery,
	)
}
