package customerhandler

import (
	"errors"
	"time"

	"github.com/overusevery/golang-echo-practice2/src/domain/entity"
	openapi "github.com/overusevery/golang-echo-practice2/src/handler/openapigenmodel/go"
	"github.com/overusevery/golang-echo-practice2/src/shared/message"
	"github.com/overusevery/golang-echo-practice2/src/shared/util"
)

type CreateCustomerResponse openapi.Customer
type CreateCustomerErrorResponse openapi.Error
type CreateCustomerMultiErrorResponse openapi.MultipleErrorResponse

func convertToCreateCustomerErrorResponse(err error) CreateCustomerMultiErrorResponse {
	messages := []openapi.ErrorElement{}
	var errList *util.ValidationErrorList
	if errors.As(err, &errList) {
		for _, err := range errList.Unwrap() {
			if customErr, ok := err.(*message.ErrorWithId); ok {
				messages = append(messages,
					openapi.ErrorElement{
						Id:  customErr.ErrorID(),
						Msg: customErr.Error(),
					},
				)
			}
		}
	}

	return CreateCustomerMultiErrorResponse{
		Errors: messages,
	}
}

func convertToCreateCustomerResponse(customer entity.Customer) CreateCustomerResponse {
	return CreateCustomerResponse{
		Id:         string(customer.ID),
		Name:       customer.Name,
		Address:    customer.Address,
		Zip:        customer.ZIP,
		Phone:      customer.Phone,
		Mktsegment: customer.MarketSegment,
		Nation:     string(customer.Nation),
		Birthdate:  time.Time(customer.Birthdate),
		Version:    customer.GetVersion(),
	}
}
