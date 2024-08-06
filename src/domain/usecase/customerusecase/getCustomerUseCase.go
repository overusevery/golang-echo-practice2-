package customerusecase

import (
	"context"

	"github.com/overusevery/golang-echo-practice2/src/domain/entity"
	"github.com/overusevery/golang-echo-practice2/src/domain/repository"
)

type GetCustomerUseCase struct {
	repository repository.CustomerRepository
}

func NewGetCustomerUseCase(repository repository.CustomerRepository) *GetCustomerUseCase {
	return &GetCustomerUseCase{
		repository: repository,
	}
}

func (uc *GetCustomerUseCase) Execute(ctx context.Context, id string) (*entity.Customer, error) {
	res, err := uc.repository.GetCustomer(ctx, id)
	return res, err
}
