package entity

import (
	"time"

	"github.com/overusevery/golang-echo-practice2/src/domain/value"
	"github.com/overusevery/golang-echo-practice2/src/shared/util"
)

type Customer struct {
	ID            int
	Name          string
	Address       string
	ZIP           string
	Phone         string
	MarketSegment string
	Nation        value.Nation
	Birthdate     value.Birthdate
}

func NewCustomer(id int, name, address, zip, phone, marketSegment, nation string, birthdate time.Time) (*Customer, util.ErrorList) {
	errList := util.NewErrorList()
	c := &Customer{
		ID:            id,
		Name:          name,
		Address:       address,
		ZIP:           zip,
		Phone:         phone,
		MarketSegment: marketSegment,
		Birthdate:     WrapNew(value.NewBirthdate, &errList)(value.NewBirthdateInput{T: birthdate, Now: time.Now()}),
		Nation:        WrapNew(value.NewNation, &errList)(nation),
	}
	if errList != nil {
		return nil, errList
	}
	return c, nil
}

type newEntitiyFun[I any, E any] func(input I) (E, util.ErrorList)

func WrapNew[I any, E any](new newEntitiyFun[I, E], errorList *util.ErrorList) func(input I) E {
	return func(input I) E {
		entity, validationErrorList := new(input)
		*errorList = errorList.Concatenate(&validationErrorList)
		return entity
	}
}
