package entity

import (
	"time"

	"github.com/overusevery/golang-echo-practice2/src/domain/entity/entityutil"
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
		Birthdate:     entityutil.WrapNew(value.NewBirthdate, &errList)(value.NewBirthdateInput{T: birthdate, Now: time.Now()}),
		Nation:        entityutil.WrapNew(value.NewNation, &errList)(nation),
	}
	if errList != nil {
		return nil, errList
	}
	return c, nil
}
