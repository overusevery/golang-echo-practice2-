package accesscontrol

import (
	"context"
	"testing"
)

func TestAccessControl_IsAllowed(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name       string
		allowlist  []string
		inputScope context.Context
		exptected  bool
	}{
		// TODO: Add test cases.
		{
			name:       "single-single",
			allowlist:  []string{"permissionA"},
			inputScope: contextWithScope("permissionA"),
			exptected:  true,
		},
		{
			name:       "single-single",
			allowlist:  []string{"permissionA"},
			inputScope: contextWithScope("permissionX"),
			exptected:  false,
		},
		{
			name:       "multi-single",
			allowlist:  []string{"permissionA", "permissionB"},
			inputScope: contextWithScope("permissionA"),
			exptected:  true,
		},
		{
			name:       "multi-single",
			allowlist:  []string{"permissionA", "permissionB"},
			inputScope: contextWithScope("permissionB"),
			exptected:  true,
		},
		{
			name:       "multi-single",
			allowlist:  []string{"permissionA", "permissionB"},
			inputScope: contextWithScope("permissionX"),
			exptected:  false,
		},
		{
			name:       "single-multi",
			allowlist:  []string{"permissionA"},
			inputScope: contextWithScope("permissionA", "permissionB"),
			exptected:  true,
		},
		{
			name:       "single-multi",
			allowlist:  []string{"permissionB"},
			inputScope: contextWithScope("permissionA", "permissionB"),
			exptected:  true,
		},
		{
			name:       "single-multi",
			allowlist:  []string{"permissionX"},
			inputScope: contextWithScope("permissionA", "permissionB"),
			exptected:  false,
		},
		{
			name:       "multi-multi",
			allowlist:  []string{"permissionA", "permissionB"},
			inputScope: contextWithScope("permissionA", "permissionB"),
			exptected:  true,
		},
		{
			name:       "multi-multi",
			allowlist:  []string{"permissionA", "permissionB"},
			inputScope: contextWithScope("permissionA", "permissionB"),
			exptected:  true,
		},
		{
			name:       "multi-multi",
			allowlist:  []string{"permissionX", "permissionA"},
			inputScope: contextWithScope("permissionA", "permissionB"),
			exptected:  true,
		},
		{
			name:       "multi-multi",
			allowlist:  []string{"permissionA", "permissionX"},
			inputScope: contextWithScope("permissionA", "permissionB"),
			exptected:  true,
		},
		{
			name:       "multi-multi",
			allowlist:  []string{"permissionB", "permissionX"},
			inputScope: contextWithScope("permissionA", "permissionB"),
			exptected:  true,
		},
		{
			name:       "multi-multi",
			allowlist:  []string{"permissionX", "permissionB"},
			inputScope: contextWithScope("permissionA", "permissionB"),
			exptected:  true,
		},
		{
			name:       "multi-multi",
			allowlist:  []string{"permissionA", "permissionB"},
			inputScope: contextWithScope("permissionA", "permissionX"),
			exptected:  true,
		},
		{
			name:       "multi-multi",
			allowlist:  []string{"permissionA", "permissionB"},
			inputScope: contextWithScope("permissionX", "permissionA"),
			exptected:  true,
		},
		{
			name:       "multi-multi",
			allowlist:  []string{"permissionA", "permissionB"},
			inputScope: contextWithScope("permissionX", "permissionB"),
			exptected:  true,
		},
		{
			name:       "multi-multi",
			allowlist:  []string{"permissionA", "permissionB"},
			inputScope: contextWithScope("permissionB", "permissionX"),
			exptected:  true,
		},
		{
			name:       "multi-multi",
			allowlist:  []string{"permissionA", "permissionB"},
			inputScope: contextWithScope("permissionX", "permissionY"),
			exptected:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := New(tt.allowlist...).IsAllowed(tt.inputScope)
			if actual != tt.exptected {
				t.Errorf("allowlist = %v, input =%v,\n actual returned value is %v, but want %v", tt.allowlist, tt.inputScope.Value("scope"), actual, tt.exptected)
			}
		})
	}
}

func contextWithScope(scopes ...string) context.Context {
	return context.WithValue(context.Background(), "scope", scopes)
}
