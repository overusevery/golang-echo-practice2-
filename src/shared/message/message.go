package message

import "github.com/overusevery/golang-echo-practice2/src/shared/util"

var (
	ERRID00001 = util.New("ERRID00001", "birthdate is too old")
	ERRID00002 = util.New("ERRID00002", "birthdate cannot be future")
	ERRID00003 = util.New("ERRID00003", "unknown country value")
	ERRID00004 = util.New("ERRID00004", "invalid request")
)
