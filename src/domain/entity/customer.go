package entity

import (
	"time"

	"github.com/overusevery/golang-echo-practice2/src/domain/value"
)

type Customer struct {
	ID            int
	Name          string
	Address       string
	ZIP           string
	Phone         string
	MarketSegment string
	Nation        string
	Birthdate     value.Birthdate
}

func NewCustomer(id int, name, address, zip, phone, marketSegment, nation string, birthdate time.Time) (*Customer, error) {
	b, err := value.NewBirthdate(birthdate, time.Now())
	if err != nil {
		return nil, err
	}
	return &Customer{
		ID:            id,
		Name:          name,
		Address:       address,
		ZIP:           zip,
		Phone:         phone,
		MarketSegment: marketSegment,
		Nation:        nation,
		Birthdate:     b,
	}, nil
}
