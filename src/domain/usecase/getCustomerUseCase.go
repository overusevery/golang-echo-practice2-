package usecase

import (
	"github.com/overusevery/golang-echo-practice2/src/domain/entity"
	"github.com/overusevery/golang-echo-practice2/src/domain/repository"
)

type GetCustomerUseCase struct {
	repository repository.CustomerRepository
}

func (uc *GetCustomerUseCase) Execute(id int) (entity.Customer, error) {
	res := uc.repository.GetCustomer(id)
	return res, nil
}
