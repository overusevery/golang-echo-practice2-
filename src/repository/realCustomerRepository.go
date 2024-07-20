package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/overusevery/golang-echo-practice2/src/domain/entity"
)

type RealCustomerRepository struct {
	db *sql.DB
}

func NewRealCustomerRepository(db *sql.DB) *RealCustomerRepository {
	return &RealCustomerRepository{
		db: db,
	}
}

func (r *RealCustomerRepository) GetCustomer(ctx context.Context, id int) entity.Customer {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	row := tx.QueryRowContext(ctx, `SELECT id, name, address, zip, phone, mktsegment, nation, birthdate FROM customers WHERE id = $1`, id)
	customer := entity.Customer{}
	err = row.Scan(&customer.ID,
		&customer.Name,
		&customer.Address,
		&customer.ZIP,
		&customer.Phone,
		&customer.MarketSegment,
		&customer.Nation,
		&customer.Birthdate,
	)
	if err != nil {
		_ = tx.Rollback()
		log.Fatal(err)
	}
	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}
	return customer
}

func (r *RealCustomerRepository) CreateCustomer(customer entity.Customer) error {
	return nil
}
