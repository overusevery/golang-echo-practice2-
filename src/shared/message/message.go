package message

var (
	ERRID00001 = New("ERRID00001", INVALID_INPUT, "birthdate is too old")
	ERRID00002 = New("ERRID00002", INVALID_INPUT, "birthdate cannot be future")
	ERRID00003 = New("ERRID00003", INVALID_INPUT, "unknown country value")
	ERRID00004 = New("ERRID00004", INVALID_INPUT, "version must be > 0")
	ERRID00005 = New("ERRID00005", DATA_NOT_FOUND, "CUSTOMER NOT FOUND")
	ERRID00006 = New("ERRID00006", CONFLICT, "Conflict")
	ERRID00007 = New("ERRID00007", INVALID_INPUT, "input value is invalid for Create Customer")
)

const (
	DATA_NOT_FOUND = iota
	INVALID_INPUT
	CONFLICT
)

type ErrorWithId struct {
	id        string
	errorType int
	msg       string
}

func New(id string, errorType int, msg string) *ErrorWithId {
	return &ErrorWithId{
		id:        id,
		errorType: errorType,
		msg:       msg,
	}

}

func (e ErrorWithId) Error() string {
	return e.msg
}

func (e ErrorWithId) ErrorID() string {
	return e.id
}
