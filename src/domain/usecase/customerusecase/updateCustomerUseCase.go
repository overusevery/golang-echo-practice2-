package customerusecase

import "errors"

type UpdateCustomerUseCase struct {
}

func (uc *UpdateCustomerUseCase) Execute(id string) error {
	//fake implementation
	if id == "2" {
		return errors.New("fake error")
	}
	return nil
}
