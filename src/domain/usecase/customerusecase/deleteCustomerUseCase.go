package customerusecase

import (
	"context"
	"errors"

	"github.com/overusevery/golang-echo-practice2/src/domain/repository"
	accesscontrol "github.com/overusevery/golang-echo-practice2/src/domain/usecase/accessControl"
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
	if accesscontrol.New("mybackendapi/editcustomer").IsNotAllowed(ctx) {
		return errors.New("not enough scope")
	}
	return uc.execute(ctx, id)
}

func (uc *DeleteCustomerUseCase) execute(ctx context.Context, id string) error {
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
