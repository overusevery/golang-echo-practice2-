package util

import (
	"testing"
)

func TestErrorWithId_Error(t *testing.T) {
	type fields struct {
		id  string
		msg string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{name: "error sholud return id and msg", fields: fields{id: "ErrorID1", msg: "Some message"}, want: "Some message"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &ErrorWithId{
				id:  tt.fields.id,
				msg: tt.fields.msg,
			}
			if got := e.Error(); got != tt.want {
				t.Errorf("ErrorWithId.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestErrorWithId_ErrorID(t *testing.T) {
	type fields struct {
		id  string
		msg string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{name: "error id should be returned", fields: fields{id: "ErrorID1", msg: "Some Message"}, want: "ErrorID1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &ErrorWithId{
				id:  tt.fields.id,
				msg: tt.fields.msg,
			}
			if got := e.ErrorID(); got != tt.want {
				t.Errorf("ErrorWithId.ErrorID() = %v, want %v", got, tt.want)
			}
		})
	}
}
