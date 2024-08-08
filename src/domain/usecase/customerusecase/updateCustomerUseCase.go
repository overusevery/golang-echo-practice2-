package customerusecase

import (
	"errors"

	"github.com/overusevery/golang-echo-practice2/src/domain/repository"
)

type UpdateCustomerUseCase struct {
	Repository repository.CustomerRepository
}

func (uc *UpdateCustomerUseCase) Execute(id string) error {
	//fake implementation
	if id == "2" {
		return errors.New("fake error")
	}
	return nil
}
