package repository

import (
	"context"

	"github.com/overusevery/golang-echo-practice2/src/domain/entity"
	"github.com/overusevery/golang-echo-practice2/src/domain/value"
	"github.com/overusevery/golang-echo-practice2/src/shared/message"
)

//go:generate mockgen -source=customerRepository.go -destination=../../repository/mock/mockCustomerRepository.go
type CustomerRepository interface {
	GetCustomer(ctx context.Context, id value.ID) (*entity.Customer, error)
	CreateCustomer(ctx context.Context, customer entity.Customer) (*entity.Customer, error)
	UpdateCustomer(ctx context.Context, customer entity.Customer) (*entity.Customer, error)
	DeleteCustomer(ctx context.Context, customer entity.DeletedCustomer) error
}

var (
	ErrCustomerNotFound = message.ERRID00005 //for GetCustomer
	ErrConflict         = message.ERRID00006
)
