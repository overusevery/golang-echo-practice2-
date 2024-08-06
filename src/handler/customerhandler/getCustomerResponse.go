package customerhandler

import (
	"time"

	"github.com/overusevery/golang-echo-practice2/src/domain/entity"
	openapi "github.com/overusevery/golang-echo-practice2/src/handler/openapigenmodel/go"
)

type GetCustomerResponse openapi.Customer
type GetCustomerErrorResponse openapi.Error
type GetCustomerMultiErrorResponse openapi.MultipleErrorResponse

func convertFrom(customer entity.Customer) GetCustomerResponse {
	return GetCustomerResponse{
		Id:         string(customer.ID),
		Name:       customer.Name,
		Address:    customer.Address,
		Zip:        customer.ZIP,
		Phone:      customer.Phone,
		Mktsegment: customer.MarketSegment,
		Nation:     string(customer.Nation),
		Birthdate:  time.Time(customer.Birthdate),
	}
}
