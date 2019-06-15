package word_test

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/matsuyoshi30/dojo5/kadai3-1/matsuyoshi30/word"
)

const TESTDIR = "./testdata"

var (
	normal = filepath.Join(TESTDIR, "test.txt")
	nofile = filepath.Join(TESTDIR, "test_nofile.txt")
	noword = filepath.Join(TESTDIR, "test_noword.txt")
)

func TestGenerateSource(t *testing.T) {
	testcases := []struct {
		name   string
		input  string
		output []string
		err    string
	}{
		{"normal", normal, []string{"apple", "banana", "coconuts"}, ""},
		{"nofile", nofile, nil, fmt.Sprintf("open %s: no such file or directory", nofile)},
		{"noword", noword, nil, "no word for typing"},
	}

	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := word.GenerateSource(tt.input)
			if err != nil {
				testGenerateSource_err(t, tt.err, err)
			}
			testGenerateSource(t, tt.output, actual)
		})
	}
}

func testGenerateSource(t *testing.T, expected []string, actual []string) {
	t.Helper()
	if len(expected) != len(actual) {
		t.Fatal("wrong length of slice")
	}

	for idx, _ := range expected {
		if expected[idx] != actual[idx] {
			t.Fatalf("expected %v, but got %v\n", expected[idx], actual[idx])
		}
	}
}

func testGenerateSource_err(t *testing.T, expected string, err error) {
	t.Helper()
	if expected != err.Error() {
		t.Fatalf("expected %v, but got %v\n", expected, err)
	}
}
