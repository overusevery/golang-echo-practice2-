package customerusecase

import (
	"context"

	"github.com/overusevery/golang-echo-practice2/src/domain/repository"
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
	_, err := uc.repository.GetCustomer(ctx, id)
	if err != nil {
		return err
	}
	err = uc.repository.DeleteCustomer(ctx, id)
	return err
}
