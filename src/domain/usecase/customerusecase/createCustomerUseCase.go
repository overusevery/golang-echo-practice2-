package customerusecase

import (
	"context"

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

func (uc *CreateCustomerUseCase) Execute(ctx context.Context, customer entity.Customer) error {
	err := uc.repository.CreateCustomer(ctx, customer)
	return err
}
