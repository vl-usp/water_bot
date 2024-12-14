package transaction

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	"github.com/vl-usp/water_bot/pkg/client/db"
	"github.com/vl-usp/water_bot/pkg/client/db/pg"
)

type manager struct {
	db db.Transactor
}

// NewTransactionManager creates a new transaction manager.
func NewTransactionManager(db db.Transactor) db.TxManager {
	return &manager{
		db: db,
	}
}

// transaction is a general function that executes the specified user handler in a transaction.
func (m *manager) transaction(ctx context.Context, opts pgx.TxOptions, fn db.Handler) (err error) {
	// If this is a nested transaction, skip the initialization of a new transaction and execute the handler.
	tx, ok := ctx.Value(pg.TxKey).(pgx.Tx)
	if ok {
		return fn(ctx)
	}

	// Start a new transaction.
	tx, err = m.db.BeginTx(ctx, opts)
	if err != nil {
		return errors.Wrap(err, "can't begin transaction")
	}

	// Make a context with a transaction.
	ctx = pg.MakeContextTx(ctx, tx)

	// Set up a deferred function to rollback or commit the transaction.
	defer func() {
		// recover from panic
		if r := recover(); r != nil {
			err = errors.Errorf("panic recovered: %v", r)
		}

		// rollback transaction, if an error occurred
		if err != nil {
			if errRollback := tx.Rollback(ctx); errRollback != nil {
				err = errors.Wrapf(err, "errRollback: %v", errRollback)
			}

			return
		}

		// commit transaction
		if nil == err {
			err = tx.Commit(ctx)
			if err != nil {
				err = errors.Wrap(err, "tx commit failed")
			}
		}
	}()

	// Execute the code inside the transaction.
	// If the function fails, return the error and the transaction is rolled back
	// or committed, depending on the result.
	if err = fn(ctx); err != nil {
		err = errors.Wrap(err, "failed executing code inside transaction")
	}

	return err
}

// ReadCommitted execute transaction with ReadCommitted isolation level
func (m *manager) ReadCommitted(ctx context.Context, f db.Handler) error {
	txOpts := pgx.TxOptions{IsoLevel: pgx.ReadCommitted}
	return m.transaction(ctx, txOpts, f)
}
