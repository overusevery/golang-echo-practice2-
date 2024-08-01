package util

import "fmt"

type ErrorWithId struct {
	id  string
	msg string
}

func New(id string, msg string) *ErrorWithId {
	return &ErrorWithId{
		id:  id,
		msg: msg,
	}

}

func (e *ErrorWithId) Error() string {
	return fmt.Sprintf("%s:%s", e.id, e.msg)
}

func (e *ErrorWithId) ErrorID() string {
	return e.id
}
