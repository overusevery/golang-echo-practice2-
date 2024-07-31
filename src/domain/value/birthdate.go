package value

import (
	"errors"
	"time"
)

type Birthdate time.Time

var (
	ErrTooOldDate = errors.New("Birthdate is too old")
)

func NewBirthdate(t time.Time) (Birthdate, error) {
	b := Birthdate(t)
	validateionErrors := b.validate()
	if validateionErrors != nil {
		return Birthdate{}, validateionErrors
	}
	return b, nil
}

func (b *Birthdate) validate() error {
	var validateionErrors error
	if time.Time(*b).Before(time.Date(1800, 1, 1, 0, 0, 0, 0, time.Local)) {
		validateionErrors = errors.Join(validateionErrors, ErrTooOldDate)
	}
	return validateionErrors

}
