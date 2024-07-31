package value

import (
	"reflect"
	"testing"
	"time"
)

func TestNewBirthdate(t *testing.T) {
	NOW := time.Now()
	type args struct {
		in0 time.Time
		now time.Time
	}
	tests := []struct {
		name    string
		args    args
		want    Birthdate
		wantErr bool
	}{
		{
			name:    "success",
			args:    args{in0: time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local), now: NOW},
			want:    Birthdate(time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local)),
			wantErr: false,
		},
		{
			name:    "Birthdate Should not be too old(<1800/1/1 is invalid)",
			args:    args{in0: time.Date(1799, 1, 1, 0, 0, 0, 0, time.Local), now: NOW},
			want:    Birthdate{},
			wantErr: true,
		},
		{
			name:    "Birthdate Should not be too old(1800/1/1 is valid)",
			args:    args{in0: time.Date(1800, 1, 1, 0, 0, 0, 0, time.Local), now: NOW},
			want:    Birthdate(time.Date(1800, 1, 1, 0, 0, 0, 0, time.Local)),
			wantErr: false,
		},
		{
			name:    "Birthdate Should not be too old(>1800/1/1 is valid)",
			args:    args{in0: time.Date(1800, 1, 2, 0, 0, 0, 0, time.Local), now: NOW},
			want:    Birthdate(time.Date(1800, 1, 2, 0, 0, 0, 0, time.Local)),
			wantErr: false,
		},
		{
			name:    "Birthdate Should not be future(now-1milisec is valid)",
			args:    args{in0: NOW.Add(time.Microsecond * -1), now: NOW},
			want:    Birthdate(NOW.Add(time.Microsecond * -1)),
			wantErr: false,
		},
		{
			name:    "Birthdate Should not be future(now is valid)",
			args:    args{in0: NOW, now: NOW},
			want:    Birthdate(NOW),
			wantErr: false,
		},
		{
			name:    "Birthdate Should not be future(now+1milisec is invalid)",
			args:    args{in0: NOW.Add(time.Microsecond), now: NOW},
			want:    Birthdate{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewBirthdate(tt.args.in0, tt.args.now)
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
