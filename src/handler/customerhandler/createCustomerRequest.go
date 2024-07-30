package customerhandler

import (
	"github.com/overusevery/golang-echo-practice2/src/domain/entity"
	openapi "github.com/overusevery/golang-echo-practice2/src/handler/openapigenmodel/go"
)

type CreateCustomerRequest openapi.NewCustomer

func (r *CreateCustomerRequest) ConvertFrom() entity.Customer {
	return entity.Customer{
		Name:          r.Name,
		Address:       r.Address,
		ZIP:           r.Zip,
		Phone:         r.Phone,
		MarketSegment: r.Mktsegment,
		Nation:        r.Nation,
		Birthdate:     r.Birthdate,
	}
}
