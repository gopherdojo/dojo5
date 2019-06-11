package main

import (
	"testing"
)

func TestOptPath(t *testing.T) {
	tests := []struct {
		input    []string
		expected []string
	}{
		{
			[]string{""},
			[]string{""},
		},
		{
			[]string{"./test"},
			[]string{"./test"},
		},
		{
			[]string{"./test", "./test/test1"},
			[]string{"./test"},
		},
		{
			[]string{"./test/test1", "./test/test2"},
			[]string{"./test/test1", "./test/test2"},
		},
	}

	for _, test := range tests {
		output := OptPath(test.input)
		if len(output) != len(test.expected) {
			t.Errorf("Failed")
		} else {
			for idx, e := range test.expected {
				if output[idx] != e {
					t.Errorf("Failed")
				}
			}
		}
	}
}
