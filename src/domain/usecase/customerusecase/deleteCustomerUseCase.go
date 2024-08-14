package customerusecase

import (
	"context"

	"github.com/overusevery/golang-echo-practice2/src/domain/repository"
	"github.com/overusevery/golang-echo-practice2/src/domain/value"
)

type DeleteCustomerUseCase struct {
	repository repository.CustomerRepository
}

func NewDeleteCustomerUseCase(repository repository.CustomerRepository) *DeleteCustomerUseCase {
	return &DeleteCustomerUseCase{
		repository: repository,
	}
}

func (uc *DeleteCustomerUseCase) Execute(ctx context.Context, id string) error {
	customer, err := uc.repository.GetCustomer(ctx, value.NewID(id))
	if err != nil {
		return err
	}

	deletedCustomer, err := customer.Delete()
	if err != nil {
		return err
	}

	err = uc.repository.DeleteCustomer(ctx, *deletedCustomer)
	return err
}
