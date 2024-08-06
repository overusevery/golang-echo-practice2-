package repository

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

const SOMEQUERY = "SELECT 1"

func TestRunInTransaction(t *testing.T) {
	t.Run("should be committed when f success", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()
		mock.ExpectBegin()
		mock.ExpectExec(SOMEQUERY)
		mock.ExpectCommit()
		f := func(ctx context.Context, tx *sql.Tx) error {
			tx.Exec(SOMEQUERY)
			return nil
		}

		err = RunInTransaction(context.Background(), db, f)
		assert.NoError(t, err)
	})

	t.Run("should be rolled back and return error when f fails", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()
		mock.ExpectBegin()
		mock.ExpectExec(SOMEQUERY)
		mock.ExpectRollback()
		f := func(ctx context.Context, tx *sql.Tx) error {
			tx.Exec(SOMEQUERY)
			return errors.New("some error")
		}

		err = RunInTransaction(context.Background(), db, f)
		assert.Equal(t, errors.New("some error"), err)
	})

	t.Run("should return error when commit fails", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()
		mock.ExpectBegin()
		mock.ExpectExec(SOMEQUERY)
		mock.ExpectCommit().WillReturnError(errors.New("some error"))
		f := func(ctx context.Context, tx *sql.Tx) error {
			tx.Exec(SOMEQUERY)
			return nil
		}

		err = RunInTransaction(context.Background(), db, f)
		assert.Equal(t, errors.New("some error"), err)
	})

	t.Run("should return error when begin fails", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()
		mock.ExpectBegin().WillReturnError(errors.New("some error"))
		f := func(ctx context.Context, tx *sql.Tx) error {
			tx.Exec(SOMEQUERY)
			return nil
		}

		err = RunInTransaction(context.Background(), db, f)
		assert.Equal(t, errors.New("some error"), err)
	})
}
