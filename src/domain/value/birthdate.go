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

type NewBirthdateInput struct {
	T   time.Time
	Now time.Time
}

func NewBirthdate(i NewBirthdateInput) (Birthdate, error) {
	t := i.T
	now := i.Now
	b := Birthdate(t)
	validateionErrors := b.validate(now)
	if validateionErrors != nil {
		return Birthdate{}, validateionErrors
	}
	return b, nil
}

func (b *Birthdate) validate(now time.Time) error {
	validateionErrors := []error{}
	if time.Time(*b).Before(time.Date(1800, 1, 1, 0, 0, 0, 0, time.Local)) {
		validateionErrors = append(validateionErrors, ErrTooOldDate)
	}
	if time.Time(*b).After(now) {
		validateionErrors = append(validateionErrors, ErrFutureDate)
	}

	if len(validateionErrors) > 0 {
		return util.NewValidationErrorList(validateionErrors...)
	}
	return nil
}
