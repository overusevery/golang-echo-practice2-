package value

import "time"

type Birthdate time.Time

func NewBirthdate(t time.Time) (Birthdate, error) {
	return Birthdate(t), nil
}
