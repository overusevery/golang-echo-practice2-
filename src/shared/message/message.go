package message

var (
	ERRID00001 = New("ERRID00001", "birthdate is too old")
	ERRID00002 = New("ERRID00002", "birthdate cannot be future")
	ERRID00003 = New("ERRID00003", "unknown country value")
	ERRID00004 = New("ERRID00004", "invalid request")
	ERRID00005 = New("ERRID00005", "invalid request") //for adding validation error fields info
)

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
