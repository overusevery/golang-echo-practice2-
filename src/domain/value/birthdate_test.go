package value

import (
	"reflect"
	"testing"
	"time"
)

func TestNewBirthdate(t *testing.T) {
	type args struct {
		in0 time.Time
	}
	tests := []struct {
		name    string
		args    args
		want    Birthdate
		wantErr bool
	}{
		{
			name:    "success",
			args:    args{in0: time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local)},
			want:    Birthdate(time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local)),
			wantErr: false,
		},
		{
			name:    "success",
			args:    args{in0: time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local)},
			want:    Birthdate(time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local)),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewBirthdate(tt.args.in0)
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
