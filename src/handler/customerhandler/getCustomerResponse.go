package customerhandler

import (
	"time"

	"github.com/overusevery/golang-echo-practice2/src/domain/entity"
	openapi "github.com/overusevery/golang-echo-practice2/src/handler/openapigenmodel/go"
)

type GetCustomerResponse openapi.Customer

func convertFrom(customer entity.Customer) GetCustomerResponse {
	return GetCustomerResponse{
		Id:         customer.ID,
		Name:       customer.Name,
		Address:    customer.Address,
		Zip:        customer.ZIP,
		Phone:      customer.Phone,
		Mktsegment: customer.MarketSegment,
		Nation:     customer.Nation,
		Birthdate:  time.Time(customer.Birthdate),
	}
}
