package customerusecase

import (
	"context"

	"github.com/overusevery/golang-echo-practice2/src/domain/entity"
	"github.com/overusevery/golang-echo-practice2/src/domain/repository"
	"github.com/overusevery/golang-echo-practice2/src/shared/util"
)

type CreateCustomerUseCase struct {
	repository repository.CustomerRepository
}

func NewCreateCustomerUseCase(repository repository.CustomerRepository) *CreateCustomerUseCase {
	return &CreateCustomerUseCase{
		repository: repository,
	}
}

func (uc *CreateCustomerUseCase) Execute(ctx context.Context, customer entity.Customer) (*entity.Customer, util.ErrorList) {
	createdCustomer, err := uc.repository.CreateCustomer(ctx, customer)
	return createdCustomer, err
}
