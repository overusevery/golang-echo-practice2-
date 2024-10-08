package value

import (
	"github.com/overusevery/golang-echo-practice2/src/shared/message"
)

type Nation string

var (
	ErrUnknownCountyValue = message.ERRID00003
)

func NewNation(s string) (Nation, error) {
	n := Nation(s)
	validateionErrors := n.validate()
	if validateionErrors != nil {
		return Nation(""), validateionErrors
	}
	return n, nil
}

var NationsList = []string{
	"JP",
	"日本",
}

func (n *Nation) validate() error {
	for _, item := range NationsList {
		if item == string(*n) {
			return nil
		}
	}
	return ErrUnknownCountyValue
}
