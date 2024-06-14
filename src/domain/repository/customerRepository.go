package repository

import "github.com/overusevery/golang-echo-practice2/src/domain/entity"

//go:generate mockgen -source=customerRepository.go -destination=../../repository/mock/mockCustomerRepository.go
type CustomerRepository interface {
	GetCustomer(id int) entity.Customer
}
