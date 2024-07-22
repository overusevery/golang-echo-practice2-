package repository

import (
	"context"
	"errors"

	"github.com/overusevery/golang-echo-practice2/src/domain/entity"
)

//go:generate mockgen -source=customerRepository.go -destination=../../repository/mock/mockCustomerRepository.go
type CustomerRepository interface {
	GetCustomer(ctx context.Context, id int) (*entity.Customer, error)
	CreateCustomer(customer entity.Customer) error
}

var (
	ErrCustomerNotFound = errors.New("CUSTOMER NOT FOUND")
)
