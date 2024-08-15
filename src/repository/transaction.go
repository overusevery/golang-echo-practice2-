package repository

import (
	"context"
	"database/sql"
)

// runInTransaction runs a function within a database transaction.
// It takes a context, a database connection, and a function to execute within the transaction.
// If the function returns an error, the transaction is rolled back. Otherwise, the transaction is committed.
func RunInTransaction(ctx context.Context, db *sql.DB, f func(ctx context.Context, tx *sql.Tx) error) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	err = f(ctx, tx)
	if err != nil {
		tx.Rollback() //nolint:errcheck
		return err
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}
