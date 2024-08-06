package customerhandler

import (
	"errors"

	openapi "github.com/overusevery/golang-echo-practice2/src/handler/openapigenmodel/go"
	"github.com/overusevery/golang-echo-practice2/src/shared/message"
	"github.com/overusevery/golang-echo-practice2/src/shared/util"
)

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

	return CreateCustomerMultiErrorResponse{
		Errors: messages,
	}
}
