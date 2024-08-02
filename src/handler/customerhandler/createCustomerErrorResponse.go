package customerhandler

import (
	"errors"

	openapi "github.com/overusevery/golang-echo-practice2/src/handler/openapigenmodel/go"
	"github.com/overusevery/golang-echo-practice2/src/shared/message"
	"github.com/overusevery/golang-echo-practice2/src/shared/util"
)

type CreateCustomerErrorResponse openapi.MultipleErrorResponse

func convertToCreateCustomerErrorResponse(errList util.ErrorList) CreateCustomerErrorResponse {
	messages := []openapi.ErrorElement{}
	for _, err := range errList {
		if customErr, ok := err.(*util.ErrorWithId); ok {
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

	return CreateCustomerErrorResponse{
		Errors: messages,
	}
}
