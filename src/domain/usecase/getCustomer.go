package usecase

import (
	"github.com/overusevery/golang-echo-practice2/src/domain/entity"
)

func GetCustomer(id int) (entity.Customer, error) {
	res := entity.Customer{ID: id}
	return res, nil
}
