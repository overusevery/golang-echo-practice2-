package util

import (
	"errors"
	"fmt"

	"github.com/overusevery/golang-echo-practice2/src/shared/message"
)

var (
	ErrValidation = message.ERRID00005
)

type ValidationErrorList struct {
	mainErr     error
	chilrenErrs []error
}

func NewValidationErrorList(err ...error) error {
	list := flattenErrorIfValitionErrorList(err...)
	vel := ValidationErrorList{
		mainErr:     ErrValidation,
		chilrenErrs: list,
	}
	return &vel
}

func flattenErrorIfValitionErrorList(err ...error) []error {
	l := []error{}
	for _, v := range err {
		var innterErrorList *ValidationErrorList
		if errors.As(v, &innterErrorList) {
			l = append(l, innterErrorList.ChilrenErrrList()...)
		} else {
			l = append(l, v)
		}
	}
	return l

}

func (e ValidationErrorList) Error() string {
	return fmt.Sprintf("validation error:%v", errors.Join(e.chilrenErrs...).Error())
}
func (e *ValidationErrorList) Unwrap() []error {
	return []error{e.mainErr, errors.Join(e.chilrenErrs...)}
}

func (e *ValidationErrorList) IsNotEmpty() bool {
	return len(e.chilrenErrs) != 0
}

func (e *ValidationErrorList) ChilrenErrrList() []error {
	return e.chilrenErrs
}
