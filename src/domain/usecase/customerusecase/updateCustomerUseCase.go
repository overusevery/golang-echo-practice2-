package customerusecase

import (
	"context"

	"github.com/overusevery/golang-echo-practice2/src/domain/entity"
	"github.com/overusevery/golang-echo-practice2/src/domain/repository"
)

type UpdateCustomerUseCase struct {
	Repository repository.CustomerRepository
}

func (uc *UpdateCustomerUseCase) Execute(ctx context.Context, id string, customer entity.Customer) (entity.Customer, error) {
	customerRes, err := uc.Repository.UpdateCustomer(ctx, customer)
	return *customerRes, err
}
