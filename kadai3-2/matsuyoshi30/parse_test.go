package godl

import (
	"errors"
	"testing"
)

func TestParseFlag(t *testing.T) {
	testcases := []struct {
		name        string
		input       []string
		expectedOpt *Option
		expectedErr error
	}{
		{"normal", []string{"./godl", "-url", "https://example.com"},
			&Option{Rt: 2, URL: "https://example.com", Output: "output"}, nil},
		{"no URL", []string{"./godl"},
			nil, errors.New("set URL")},
	}

	for _, tt := range testcases {
		_, err := ParseFlag(tt.input[1:]...)
		if err != nil {
			if tt.expectedErr == nil {
				t.Fatalf("unexpected err")
			}
		}
	}
}
