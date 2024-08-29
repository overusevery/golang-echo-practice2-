package customerusecase

import (
	"context"
	"errors"
	"time"

	"github.com/overusevery/golang-echo-practice2/src/domain/entity"
	"github.com/overusevery/golang-echo-practice2/src/domain/repository"
	accesscontrol "github.com/overusevery/golang-echo-practice2/src/domain/usecase/accessControl"
	"github.com/overusevery/golang-echo-practice2/src/domain/value"
)

type UpdateCustomerUseCase struct {
	repository repository.CustomerRepository
}

func NewUpdateCustomerUseCase(repository repository.CustomerRepository) *UpdateCustomerUseCase {
	return &UpdateCustomerUseCase{
		repository: repository,
	}
}

func (uc *UpdateCustomerUseCase) Execute(
	ctx context.Context,
	id string,
	input UpdateCustomerUseCaseInput,
) (*entity.Customer, error) {
	if accesscontrol.New("mybackendapi/editcustomer").IsNotAllowed(ctx) {
		return nil, errors.New("not enough scope")
	}
	return uc.execute(ctx, id, input)
}

func (uc *UpdateCustomerUseCase) execute(
	ctx context.Context,
	id string,
	input UpdateCustomerUseCaseInput,
) (*entity.Customer, error) {
	currentCustomer, err := uc.repository.GetCustomer(ctx, value.NewID(id))
	if err != nil {
		return nil, err
	}
	customer, err := currentCustomer.ChangeInfo(
		input.Name,
		input.Address,
		input.ZIP,
		input.Phone,
		input.MarketSegment,
		input.Nation,
		input.Birthdate,
	)
	if err != nil {
		return nil, err
	}
	customerRes, err := uc.repository.UpdateCustomer(ctx, *customer)
	if err != nil {
		return nil, err
	}
	return customerRes, nil
}

type UpdateCustomerUseCaseInput struct {
	Name          string
	Address       string
	ZIP           string
	Phone         string
	MarketSegment string
	Nation        string
	Birthdate     time.Time
	Version       int
}
