package customerusecase

import (
	"github.com/overusevery/golang-echo-practice2/src/domain/entity"
	"github.com/overusevery/golang-echo-practice2/src/domain/repository"
)

type CreateCustomerUseCase struct {
	repository repository.CustomerRepository
}

func NewCreateCustomerUseCase(repository repository.CustomerRepository) *CreateCustomerUseCase {
	return &CreateCustomerUseCase{
		repository: repository,
	}
}

func (uc *CreateCustomerUseCase) Execute(customer entity.Customer) error {
	err := uc.repository.CreateCustomer(customer)
	return err
}
