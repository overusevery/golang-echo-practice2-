package customerusecase

import (
	"context"
	"errors"

	"github.com/overusevery/golang-echo-practice2/src/domain/entity"
	"github.com/overusevery/golang-echo-practice2/src/domain/repository"
	accesscontrol "github.com/overusevery/golang-echo-practice2/src/domain/usecase/accessControl"
	"github.com/overusevery/golang-echo-practice2/src/domain/value"
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
	if accesscontrol.New("mybackendapi/getcustomer").IsNotAllowed(ctx) {
		return nil, errors.New("not enough scope")
	}
	return uc.execute(ctx, id)
}
func (uc *GetCustomerUseCase) execute(ctx context.Context, id string) (*entity.Customer, error) {
	res, err := uc.repository.GetCustomer(ctx, value.NewID(id))
	return res, err
}
