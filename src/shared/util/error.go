package util

type ErrorList []error

func NewErrorList(err ...error) ErrorList {
	if len(err) == 0 {
		return nil
	}
	return append(ErrorList{}, err...)
}

func (e *ErrorList) Append(err error) ErrorList {
	errorList := []error(*e)
	errorList = append(errorList, err)
	return ErrorList(errorList)
}

func (e *ErrorList) Contains(err error) bool {
	for _, v := range *e {
		if v == err {
			return true
		}
	}
	return false
}

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

func (e ErrorWithId) Error() string {
	return e.msg
}

func (e ErrorWithId) ErrorID() string {
	return e.id
}
