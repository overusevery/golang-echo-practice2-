package customerhandler

import (
	"github.com/overusevery/golang-echo-practice2/src/domain/entity"
	openapi "github.com/overusevery/golang-echo-practice2/src/handler/openapigenmodel/go"
	"github.com/overusevery/golang-echo-practice2/src/shared/util"
)

type CreateCustomerRequest openapi.NewCustomer

func (r *CreateCustomerRequest) ConvertFrom() (*entity.Customer, util.ErrorList) {
	c, errList := entity.NewCustomer(
		0,
		r.Name,
		r.Address,
		r.Zip,
		r.Phone,
		r.Mktsegment,
		r.Nation,
		r.Birthdate,
	)
	if errList != nil {
		return nil, errList
	}
	return c, nil
}
