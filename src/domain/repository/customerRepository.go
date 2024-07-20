package repository

import (
	"context"

	"github.com/overusevery/golang-echo-practice2/src/domain/entity"
)

//go:generate mockgen -source=customerRepository.go -destination=../../repository/mock/mockCustomerRepository.go
type CustomerRepository interface {
	GetCustomer(ctx context.Context, id int) entity.Customer
	CreateCustomer(customer entity.Customer) error
}
