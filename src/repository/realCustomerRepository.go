package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/overusevery/golang-echo-practice2/src/domain/entity"
	"github.com/overusevery/golang-echo-practice2/src/domain/repository"
)

type RealCustomerRepository struct {
	db *sql.DB
}

func NewRealCustomerRepository(db *sql.DB) *RealCustomerRepository {
	return &RealCustomerRepository{
		db: db,
	}
}

func (r *RealCustomerRepository) GetCustomer(ctx context.Context, id int) (*entity.Customer, error) {
	var entityCustomer *entity.Customer
	errTranscation := RunInTransaction(ctx, r.db, func(ctx context.Context, tx *sql.Tx) error {
		row := tx.QueryRowContext(ctx, `SELECT id, name, address, zip, phone, mktsegment, nation, birthdate FROM customers WHERE id = $1`, id)
		dbCustomer := DBCustomer{}
		err := row.Scan(&dbCustomer.ID,
			&dbCustomer.Name,
			&dbCustomer.Address,
			&dbCustomer.ZIP,
			&dbCustomer.Phone,
			&dbCustomer.MarketSegment,
			&dbCustomer.Nation,
			&dbCustomer.Birthdate,
		)
		if err != nil {
			switch {
			case err == sql.ErrNoRows:
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
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	var id int
	if err = tx.QueryRowContext(ctx,
		`INSERT INTO customers (name, address, zip, phone, mktsegment, nation, birthdate)
        VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id;`,
		customer.Name,
		customer.Address,
		customer.ZIP,
		customer.Phone,
		customer.MarketSegment,
		customer.Nation,
		time.Time(customer.Birthdate),
	).Scan(&id); err != nil {
		_ = tx.Rollback()
		return nil, err
	}
	row := tx.QueryRowContext(ctx, `SELECT id, name, address, zip, phone, mktsegment, nation, birthdate FROM customers WHERE id = $1`, id)
	dbCustomer := DBCustomer{}
	err = row.Scan(&dbCustomer.ID,
		&dbCustomer.Name,
		&dbCustomer.Address,
		&dbCustomer.ZIP,
		&dbCustomer.Phone,
		&dbCustomer.MarketSegment,
		&dbCustomer.Nation,
		&dbCustomer.Birthdate,
	)
	if err != nil {
		_ = tx.Rollback()
		return nil, err
	}
	if err := tx.Commit(); err != nil {
		return nil, err
	}

	entityCustomer, errList := dbCustomer.convertToEntity()
	if errList != nil {
		return nil, errList
	}
	return entityCustomer, nil
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
	)
	if err != nil {
		return nil, err
	}
	return c, nil
}
