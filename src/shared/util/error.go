package util

import (
	"errors"
)

type ValidationErrorList struct {
	errors []error
}

func NewValidationErrorList(err ...error) error {
	if err == nil { // if no err
		return nil
	}
	list := flattenErrorIfValitionErrorList(err...)
	vel := ValidationErrorList{
		errors: list,
	}
	return &vel
}

func flattenErrorIfValitionErrorList(err ...error) []error {
	l := []error{}
	for _, v := range err {
		var innterErrorList *ValidationErrorList
		if errors.As(v, &innterErrorList) {
			l = append(l, innterErrorList.Unwrap()...)
		} else {
			l = append(l, v)
		}
	}
	return l

}

func (e ValidationErrorList) Error() string {
	return errors.Join(e.Unwrap()...).Error()
}
func (e *ValidationErrorList) Unwrap() []error {
	return e.errors
}
