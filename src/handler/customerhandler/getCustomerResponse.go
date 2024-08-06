package customerhandler

import (
	"errors"
	"time"

	"github.com/overusevery/golang-echo-practice2/src/domain/entity"
	openapi "github.com/overusevery/golang-echo-practice2/src/handler/openapigenmodel/go"
	"github.com/overusevery/golang-echo-practice2/src/shared/message"
	"github.com/overusevery/golang-echo-practice2/src/shared/util"
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

func convertToGetCustomerMultiResponse(err error) GetCustomerMultiErrorResponse {
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
			} else {
				messages = append(messages,
					openapi.ErrorElement{
						Id:  message.ERRID00004.ErrorID(),
						Msg: errors.Join(message.ERRID00004, err).Error(),
					},
				)
			}
		}
	}

	return GetCustomerMultiErrorResponse{
		Errors: messages,
	}
}
