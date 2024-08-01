package customerhandler

import (
	"fmt"

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
			// Handle MyCustomError specifically
			fmt.Println("Custom error:", customErr.Error())
			messages = append(messages,
				ErrorMsg{
					ID:  customErr.ErrorID(),
					Msg: customErr.Error(),
				},
			)
		} else {
			// Handle generic error
			fmt.Println("Generic error:", err.Error())
		}

	}

	return CreateCustomerErrorResponse{
		ErrorMsgs: messages,
	}
}
