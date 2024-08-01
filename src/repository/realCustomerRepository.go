package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/overusevery/golang-echo-practice2/src/domain/entity"
	"github.com/overusevery/golang-echo-practice2/src/domain/repository"
	"github.com/overusevery/golang-echo-practice2/src/shared/util"
)

type RealCustomerRepository struct {
	db *sql.DB
}

func NewRealCustomerRepository(db *sql.DB) *RealCustomerRepository {
	return &RealCustomerRepository{
		db: db,
	}
}

func (r *RealCustomerRepository) GetCustomer(ctx context.Context, id int) (*entity.Customer, util.ErrorList) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, util.ErrorList{err}
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
	if err == sql.ErrNoRows {
		return nil, util.ErrorList{repository.ErrCustomerNotFound}
	}
	if err != nil {
		_ = tx.Rollback()
		return nil, util.ErrorList{err}
	}
	if err := tx.Commit(); err != nil {
		return nil, util.ErrorList{err}
	}
	entityCustomer, errList := dbCustomer.convertToEntity()
	if errList != nil {
		return nil, errList
	}

	return entityCustomer, nil
}

func (r *RealCustomerRepository) CreateCustomer(ctx context.Context, customer entity.Customer) (*entity.Customer, util.ErrorList) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, util.ErrorList{err}
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
		return nil, util.ErrorList{err}
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
		return nil, util.ErrorList{err}
	}
	if err := tx.Commit(); err != nil {
		return nil, util.ErrorList{err}
	}

	entityCustomer, errList := dbCustomer.convertToEntity()
	if errList != nil {
		return nil, errList
	}
	return entityCustomer, nil
}

type DBCustomer struct {
	ID            int
	Name          string
	Address       string
	ZIP           string
	Phone         string
	MarketSegment string
	Nation        string
	Birthdate     time.Time
}

func (d *DBCustomer) convertToEntity() (*entity.Customer, util.ErrorList) {
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
