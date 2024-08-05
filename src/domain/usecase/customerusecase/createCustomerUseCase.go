package customerusecase

import (
	"context"
	"errors"
	"time"

	"github.com/overusevery/golang-echo-practice2/src/domain/entity"
	"github.com/overusevery/golang-echo-practice2/src/domain/repository"
)

var (
	ErrInvalidInputCreateCustomerUseCase = errors.New("input value is invalid for Create Customer")
)

type CreateCustomerUseCaseInput struct {
	Name          string
	Address       string
	ZIP           string
	Phone         string
	MarketSegment string
	Nation        string
	Birthdate     time.Time
}
type CreateCustomerUseCase struct {
	repository repository.CustomerRepository
}

func NewCreateCustomerUseCase(repository repository.CustomerRepository) *CreateCustomerUseCase {
	return &CreateCustomerUseCase{
		repository: repository,
	}
}

func (uc *CreateCustomerUseCase) Execute(ctx context.Context, input CreateCustomerUseCaseInput) (*entity.Customer, error) {
	customer, err := entity.NewCustomerNotRegistered(
		input.Name,
		input.Address,
		input.ZIP,
		input.Phone,
		input.MarketSegment,
		input.Nation,
		input.Birthdate,
	)
	if err != nil {
		return nil, errors.Join(ErrInvalidInputCreateCustomerUseCase, err)
	}
	createdCustomer, err := uc.repository.CreateCustomer(ctx, *customer)
	return createdCustomer, err
}
