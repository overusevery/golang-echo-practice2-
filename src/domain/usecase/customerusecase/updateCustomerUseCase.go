package customerusecase

import (
	"context"
	"errors"

	"github.com/overusevery/golang-echo-practice2/src/domain/entity"
	"github.com/overusevery/golang-echo-practice2/src/domain/repository"
)

type UpdateCustomerUseCase struct {
	Repository repository.CustomerRepository
}

func (uc *UpdateCustomerUseCase) Execute(ctx context.Context, id string) error {
	//fake implementation
	if uc.Repository != nil {
		uc.Repository.UpdateCustomer(ctx, entity.Customer{})
	}
	if id == "2" {
		return errors.New("fake error")
	}
	return nil
}
