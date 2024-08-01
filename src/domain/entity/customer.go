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
	b, errList := value.NewBirthdate(birthdate, time.Now())
	n, errListNation := value.NewNation(nation)
	errList = errList.Concatenate(&errListNation)
	if errList != nil {
		return nil, errList
	}
	return &Customer{
		ID:            id,
		Name:          name,
		Address:       address,
		ZIP:           zip,
		Phone:         phone,
		MarketSegment: marketSegment,
		Nation:        n,
		Birthdate:     b,
	}, nil
}
