package entity

import (
	"errors"
	"reflect"
	"testing"
)

func TestNewAggregate(t *testing.T) {
	tests := []struct {
		name    string
		version int
		want    Aggregate
		err     error
	}{
		{name: "cannot be negative", version: -1, want: Aggregate{version: 0}, err: ErrInvalidVersion},
		{name: "cannot be zero", version: 0, want: Aggregate{version: 0}, err: ErrInvalidVersion},
		{name: "can be positive", version: 1, want: Aggregate{version: 1}, err: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewAggregate(tt.version)
			if !errors.Is(err, tt.err) {
				t.Errorf("NewAggregate() error = %v, wanted %v", err, tt.err)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAggregate() = %v, want %v", got, tt.want)
			}
		})
	}
}
