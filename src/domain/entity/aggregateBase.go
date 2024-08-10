package entity

import (
	"errors"
	"fmt"
)

var ErrInvalidVersion = errors.New("version must be > 0")

type Aggregate struct {
	version int
}

func NewAggregate(version int) (*Aggregate, error) {
	if version <= 0 {
		return nil, fmt.Errorf("version cannot be %v:%w", version, ErrInvalidVersion)
	}

	a := Aggregate{version: version}
	return &a, nil
}

func (a Aggregate) GetVersion() int {
	return a.version
}
