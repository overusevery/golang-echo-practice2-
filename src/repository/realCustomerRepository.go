package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/overusevery/golang-echo-practice2/src/domain/entity"
	"github.com/overusevery/golang-echo-practice2/src/domain/repository"
	"github.com/overusevery/golang-echo-practice2/src/domain/value"
)

type RealCustomerRepository struct {
	db *sql.DB
}

func NewRealCustomerRepository(db *sql.DB) *RealCustomerRepository {
	return &RealCustomerRepository{
		db: db,
	}
}

func (r *RealCustomerRepository) GetCustomer(ctx context.Context, id value.ID) (*entity.Customer, error) {
	var entityCustomer *entity.Customer
	errTranscation := RunInTransaction(ctx, r.db, func(ctx context.Context, tx *sql.Tx) error {
		row := tx.QueryRowContext(ctx, `SELECT id, name, address, zip, phone, mktsegment, nation, birthdate, version FROM customers WHERE id = $1`, id)
		dbCustomer := DBCustomer{}
		err := row.Scan(&dbCustomer.ID,
			&dbCustomer.Name,
			&dbCustomer.Address,
			&dbCustomer.ZIP,
			&dbCustomer.Phone,
			&dbCustomer.MarketSegment,
			&dbCustomer.Nation,
			&dbCustomer.Birthdate,
			&dbCustomer.Version,
		)
		if err != nil {
			switch {
			case errors.Is(err, sql.ErrNoRows):
				return repository.ErrCustomerNotFound
			default:
				return err
			}
		}
		entityCustomer, err = dbCustomer.convertToEntity()
		if err != nil {
			return err
		}
		return nil
	})
	return entityCustomer, errTranscation
}

func (r *RealCustomerRepository) CreateCustomer(ctx context.Context, customer entity.Customer) (*entity.Customer, error) {
	var entityCustomer *entity.Customer
	errRun := RunInTransaction(ctx, r.db, func(ctx context.Context, tx *sql.Tx) error {
		_, err := tx.ExecContext(ctx,
			`INSERT INTO customers (id, name, address, zip, phone, mktsegment, nation, birthdate, version)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`,
			customer.ID,
			customer.Name,
			customer.Address,
			customer.ZIP,
			customer.Phone,
			customer.MarketSegment,
			customer.Nation,
			time.Time(customer.Birthdate),
			customer.GetVersion(),
		)
		if err != nil {
			return err
		}

		row := tx.QueryRowContext(ctx, `SELECT id, name, address, zip, phone, mktsegment, nation, birthdate, version FROM customers WHERE id = $1`, customer.ID)
		dbCustomer := DBCustomer{}
		err = row.Scan(&dbCustomer.ID,
			&dbCustomer.Name,
			&dbCustomer.Address,
			&dbCustomer.ZIP,
			&dbCustomer.Phone,
			&dbCustomer.MarketSegment,
			&dbCustomer.Nation,
			&dbCustomer.Birthdate,
			&dbCustomer.Version,
		)
		if err != nil {
			return err
		}

		entityCustomer, err = dbCustomer.convertToEntity()
		if err != nil {
			return err
		}
		return nil
	})
	return entityCustomer, errRun
}

func (r *RealCustomerRepository) UpdateCustomer(ctx context.Context, customer entity.Customer) (*entity.Customer, error) {
	var entityCustomer *entity.Customer
	errRun := RunInTransaction(ctx, r.db, func(ctx context.Context, tx *sql.Tx) error {
		result, err := tx.ExecContext(ctx,
			`UPDATE customers SET name = $1, address = $2, zip = $3, phone = $4, mktsegment = $5, nation = $6, birthdate = $7, version = $8 WHERE id = $9 and version = $10`,
			customer.Name,
			customer.Address,
			customer.ZIP,
			customer.Phone,
			customer.MarketSegment,
			customer.Nation,
			time.Time(customer.Birthdate),
			customer.GetVersion()+1,
			customer.ID,
			customer.GetVersion(),
		)
		if err != nil {
			return err
		}
		numAffectedRows, err := result.RowsAffected()
		if err != nil {
			return err
		}
		if numAffectedRows == 0 {
			return repository.ErrConflict
		}

		row := tx.QueryRowContext(ctx, `SELECT id, name, address, zip, phone, mktsegment, nation, birthdate, version FROM customers WHERE id = $1`, customer.ID)
		dbCustomer := DBCustomer{}
		err = row.Scan(&dbCustomer.ID,
			&dbCustomer.Name,
			&dbCustomer.Address,
			&dbCustomer.ZIP,
			&dbCustomer.Phone,
			&dbCustomer.MarketSegment,
			&dbCustomer.Nation,
			&dbCustomer.Birthdate,
			&dbCustomer.Version,
		)
		if err != nil {
			return err
		}

		entityCustomer, err = dbCustomer.convertToEntity()
		if err != nil {
			return err
		}
		return nil
	})
	return entityCustomer, errRun
}

func (r *RealCustomerRepository) DeleteCustomer(ctx context.Context, customer entity.DeletedCustomer) error {
	return RunInTransaction(ctx, r.db, func(ctx context.Context, tx *sql.Tx) error {
		_, err := tx.ExecContext(ctx, `DELETE FROM customers WHERE id = $1`, customer.ID)
		if err != nil {
			return err
		}
		return nil
	})
}

type DBCustomer struct {
	ID            string
	Name          string
	Address       string
	ZIP           string
	Phone         string
	MarketSegment string
	Nation        string
	Birthdate     time.Time
	Version       int
}

func (d *DBCustomer) convertToEntity() (*entity.Customer, error) {
	c, err := entity.NewCustomer(
		d.ID,
		d.Name,
		d.Address,
		d.ZIP,
		d.Phone,
		d.MarketSegment,
		d.Nation,
		d.Birthdate,
		d.Version,
	)
	if err != nil {
		return nil, err
	}
	return c, nil
}
