package repository

import (
	"context"
	"errors"

	"github.com/overusevery/golang-echo-practice2/src/domain/entity"
	"github.com/overusevery/golang-echo-practice2/src/domain/value"
)

//go:generate mockgen -source=customerRepository.go -destination=../../repository/mock/mockCustomerRepository.go
type CustomerRepository interface {
	GetCustomer(ctx context.Context, id value.ID) (*entity.Customer, error)
	CreateCustomer(ctx context.Context, customer entity.Customer) (*entity.Customer, error)
	UpdateCustomer(ctx context.Context, customer entity.Customer) (*entity.Customer, error)
	DeleteCustomer(ctx context.Context, customer entity.DeletedCustomer) error
}

var (
	ErrCustomerNotFound = errors.New("CUSTOMER NOT FOUND") //for GetCustomer
	ErrConflict         = errors.New("Conflict")
)
