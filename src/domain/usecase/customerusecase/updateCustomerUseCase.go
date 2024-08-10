package customerusecase

import (
	"context"
	"time"

	"github.com/overusevery/golang-echo-practice2/src/domain/entity"
	"github.com/overusevery/golang-echo-practice2/src/domain/repository"
)

type UpdateCustomerUseCase struct {
	repository repository.CustomerRepository
}

func NewUpdateCustomerUseCase(repository repository.CustomerRepository) *UpdateCustomerUseCase {
	return &UpdateCustomerUseCase{
		repository: repository,
	}
}

func (uc *UpdateCustomerUseCase) Execute(ctx context.Context, id string, input UpdateCustomerUseCaseInput) (*entity.Customer, error) {
	customer, err := entity.NewCustomer(
		id,
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
}
