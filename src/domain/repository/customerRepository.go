package repository

import (
	"context"
	"errors"

	"github.com/overusevery/golang-echo-practice2/src/domain/entity"
	"github.com/overusevery/golang-echo-practice2/src/shared/util"
)

//go:generate mockgen -source=customerRepository.go -destination=../../repository/mock/mockCustomerRepository.go
type CustomerRepository interface {
	GetCustomer(ctx context.Context, id int) (*entity.Customer, util.ErrorList)
	CreateCustomer(ctx context.Context, customer entity.Customer) (*entity.Customer, util.ErrorList)
}

var (
	ErrCustomerNotFound = errors.New("CUSTOMER NOT FOUND")
)
