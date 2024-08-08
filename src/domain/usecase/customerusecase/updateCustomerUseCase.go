package customerusecase

import (
	"context"
	"time"

	"github.com/overusevery/golang-echo-practice2/src/domain/entity"
	"github.com/overusevery/golang-echo-practice2/src/domain/repository"
)

type UpdateCustomerUseCase struct {
	Repository repository.CustomerRepository
}

func (uc *UpdateCustomerUseCase) Execute(ctx context.Context, id string, input UpdateCustomerUseCaseInput) (entity.Customer, error) {
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
		return *customer, err
	}
	customerRes, err := uc.Repository.UpdateCustomer(ctx, *customer)
	return *customerRes, err
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
