package customerhandler

import (
	"github.com/overusevery/golang-echo-practice2/src/domain/entity"
	openapi "github.com/overusevery/golang-echo-practice2/src/handler/openapigenmodel/go"
)

type CreateCustomerRequest openapi.NewCustomer

func (r *CreateCustomerRequest) ConvertFrom() entity.Customer {
	c, _ := entity.NewCustomer(
		0,
		r.Name,
		r.Address,
		r.Zip,
		r.Phone,
		r.Mktsegment,
		r.Nation,
		r.Birthdate,
	)
	//ToDo:Error handling
	return *c
}
