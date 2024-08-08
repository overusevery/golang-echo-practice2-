package customerhandler

import (
	"time"

	"github.com/overusevery/golang-echo-practice2/src/domain/entity"
	openapi "github.com/overusevery/golang-echo-practice2/src/handler/openapigenmodel/go"
)

type UpdateCustomerResponse openapi.Customer

// type CreateCustomerErrorResponse openapi.Error
// type CreateCustomerMultiErrorResponse openapi.MultipleErrorResponse

func convertToUpdateCustomerResponse(customer entity.Customer) UpdateCustomerResponse {
	return UpdateCustomerResponse{
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
