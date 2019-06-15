package parse_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/matsuyoshi30/dojo5/kadai3-1/matsuyoshi30/parse"
)

func TestParse(t *testing.T) {
	testcases := []struct {
		name     string
		input    string
		expected *parse.Option
	}{
		{"noargs", "", &parse.Option{20, "default.txt"}},
		{"lt", "-lt=30", &parse.Option{30, "default.txt"}},
		{"wf", "-wf=change.txt", &parse.Option{20, "change.txt"}},
	}

	for _, tt := range testcases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			actual := parse.ParseFlag(tt.input)
			if diff := cmp.Diff(tt.expected, actual); diff != "" {
				t.Fatalf("expected %v, but got %v", tt.expected, actual)
			}
		})
	}
}
