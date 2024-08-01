package customerhandler

import (
	"errors"

	"github.com/overusevery/golang-echo-practice2/src/shared/message"
	"github.com/overusevery/golang-echo-practice2/src/shared/util"
)

type CreateCustomerErrorResponse struct {
	ErrorMsgs []ErrorMsg `json:"errors"`
}
type ErrorMsg struct {
	ID  string `json:"id"`
	Msg string `json:"msg"`
}

func convertToCreateCustomerErrorResponse(errList util.ErrorList) CreateCustomerErrorResponse {
	messages := []ErrorMsg{}
	for _, err := range errList {
		if customErr, ok := err.(*util.ErrorWithId); ok {
			messages = append(messages,
				ErrorMsg{
					ID:  customErr.ErrorID(),
					Msg: customErr.Error(),
				},
			)
		} else {
			messages = append(messages,
				ErrorMsg{
					ID:  message.ERRID00004.ErrorID(),
					Msg: errors.Join(message.ERRID00004, err).Error(),
				},
			)
		}

	}

	return CreateCustomerErrorResponse{
		ErrorMsgs: messages,
	}
}
