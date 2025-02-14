package pg

import (
	"context"
	"log/slog"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	"github.com/vl-usp/water_bot/pkg/client/db"
)

type pgClient struct {
	masterDBC db.DB
}

// New creates a new PostgreSQL client.
func New(ctx context.Context, dsn string, log *slog.Logger, debugMode bool) (db.Client, error) {
	dbc, err := pgxpool.Connect(ctx, dsn)
	if err != nil {
		return nil, errors.Errorf("failed to connect to db: %v", err)
	}

	return &pgClient{
		masterDBC: NewDB(dbc, log, debugMode),
	}, nil
}

// DB returns the master DB.
func (c *pgClient) DB() db.DB {
	return c.masterDBC
}

// Close closes the client.
func (c *pgClient) Close() error {
	if c.masterDBC != nil {
		c.masterDBC.Close()
	}

	return nil
}
