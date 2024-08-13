package entity

import (
	"time"

	"github.com/overusevery/golang-echo-practice2/src/domain/entity/entityutil"
	"github.com/overusevery/golang-echo-practice2/src/domain/value"
	"github.com/overusevery/golang-echo-practice2/src/shared/util"
)

type Customer struct {
	Aggregate
	ID            value.ID
	Name          string
	Address       string
	ZIP           string
	Phone         string
	MarketSegment string
	Nation        value.Nation
	Birthdate     value.Birthdate
}

func NewCustomer(id string, name, address, zip, phone, marketSegment, nation string, birthdate time.Time, version int) (*Customer, error) {
	errList := []error{}
	c := &Customer{
		ID:            value.NewID(id),
		Name:          name,
		Address:       address,
		ZIP:           zip,
		Phone:         phone,
		MarketSegment: marketSegment,
		Birthdate:     entityutil.WrapNew(value.NewBirthdate, &errList)(value.NewBirthdateInput{T: birthdate, Now: time.Now()}),
		Nation:        entityutil.WrapNew(value.NewNation, &errList)(nation),
		Aggregate:     entityutil.WrapNew(NewAggregate, &errList)(version),
	}
	if len(errList) > 0 {
		return nil, util.NewValidationErrorList(errList...)
	}
	return c, nil
}

func NewCustomerNotRegistered(name, address, zip, phone, marketSegment, nation string, birthdate time.Time) (*Customer, error) {
	c, err := NewCustomer(value.GenerateNewIDString(), name, address, zip, phone, marketSegment, nation, birthdate, 1)
	return c, err
}

func (c Customer) ChangeInfo(id string, name, address, zip, phone, marketSegment, nation string, birthdate time.Time, version int) (*Customer, error) {
	return NewCustomer(id, name, address, zip, phone, marketSegment, nation, birthdate, version)
}
