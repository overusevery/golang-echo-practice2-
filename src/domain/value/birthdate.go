package value

import (
	"time"

	"github.com/overusevery/golang-echo-practice2/src/shared/message"
	"github.com/overusevery/golang-echo-practice2/src/shared/util"
)

type Birthdate time.Time

var (
	ErrTooOldDate = message.ERRID00001
	ErrFutureDate = message.ERRID00002
)

func NewBirthdate(t time.Time, now time.Time) (Birthdate, util.ErrorList) {
	b := Birthdate(t)
	validateionErrors := b.validate(now)
	if validateionErrors != nil {
		return Birthdate{}, validateionErrors
	}
	return b, nil
}

func (b *Birthdate) validate(now time.Time) util.ErrorList {
	var validateionErrors util.ErrorList
	if time.Time(*b).Before(time.Date(1800, 1, 1, 0, 0, 0, 0, time.Local)) {
		validateionErrors = validateionErrors.Append(ErrTooOldDate)
	}
	if time.Time(*b).After(now) {
		validateionErrors = validateionErrors.Append(ErrFutureDate)
	}
	return validateionErrors

}
