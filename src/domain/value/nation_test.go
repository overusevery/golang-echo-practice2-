package value

import (
	"testing"

	"github.com/overusevery/golang-echo-practice2/src/shared/util"
	"github.com/stretchr/testify/assert"
)

func TestNewNation(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name  string
		args  args
		want  Nation
		want1 util.ErrorList
	}{
		{
			name:  "success",
			args:  args{s: "JP"},
			want:  Nation("JP"),
			want1: util.NewErrorList(),
		},
		{
			name: "invalid small case",
			args: args{
				s: "Jp",
			},
			want:  Nation(""),
			want1: util.NewErrorList(ErrUnknownCountyValue),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := NewNation(tt.args.s)
			if got != tt.want {
				t.Errorf("NewNation() got1 = %v, want %v", got, tt.want)
			}
			assert.Equal(t, got1, tt.want1)
		})
	}
}
