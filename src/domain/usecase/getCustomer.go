package usecase

import (
	"github.com/overusevery/golang-echo-practice2/src/domain/entity"
)

type GetCustomerUseCase struct{}

func (uc *GetCustomerUseCase) Execute(id int) (entity.Customer, error) {
	res := entity.Customer{ID: id}
	return res, nil
}
