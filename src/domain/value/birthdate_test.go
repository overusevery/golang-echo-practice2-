package value

import (
	"reflect"
	"testing"
	"time"
)

func TestNewBirthdate(t *testing.T) {
	NOW := time.Now()
	tests := []struct {
		name    string
		arg     NewBirthdateInput
		want    Birthdate
		wantErr bool
	}{
		{
			name:    "success",
			arg:     NewBirthdateInput{T: time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local), Now: NOW},
			want:    Birthdate(time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local)),
			wantErr: false,
		},
		{
			name:    "Birthdate Should not be too old(<1800/1/1 is invalid)",
			arg:     NewBirthdateInput{T: time.Date(1799, 1, 1, 0, 0, 0, 0, time.Local), Now: NOW},
			want:    Birthdate{},
			wantErr: true,
		},
		{
			name:    "Birthdate Should not be too old(1800/1/1 is valid)",
			arg:     NewBirthdateInput{T: time.Date(1800, 1, 1, 0, 0, 0, 0, time.Local), Now: NOW},
			want:    Birthdate(time.Date(1800, 1, 1, 0, 0, 0, 0, time.Local)),
			wantErr: false,
		},
		{
			name:    "Birthdate Should not be too old(>1800/1/1 is valid)",
			arg:     NewBirthdateInput{T: time.Date(1800, 1, 2, 0, 0, 0, 0, time.Local), Now: NOW},
			want:    Birthdate(time.Date(1800, 1, 2, 0, 0, 0, 0, time.Local)),
			wantErr: false,
		},
		{
			name:    "Birthdate Should not be future(now-1milisec is valid)",
			arg:     NewBirthdateInput{T: NOW.Add(time.Microsecond * -1), Now: NOW},
			want:    Birthdate(NOW.Add(time.Microsecond * -1)),
			wantErr: false,
		},
		{
			name:    "Birthdate Should not be future(now is valid)",
			arg:     NewBirthdateInput{T: NOW, Now: NOW},
			want:    Birthdate(NOW),
			wantErr: false,
		},
		{
			name:    "Birthdate Should not be future(now+1milisec is invalid)",
			arg:     NewBirthdateInput{T: NOW.Add(time.Microsecond), Now: NOW},
			want:    Birthdate{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewBirthdate(tt.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewBirthdate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBirthdate() = %v, want %v", got, tt.want)
			}
		})
	}
}
