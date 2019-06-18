package words

import (
	"reflect"
	"testing"

	"golang.org/x/xerrors"
)

func TestImport(t *testing.T) {
	t.Helper()

	runImportTests := []struct {
		testCase string
		path     string
		out      []string
		err      error
	}{
		{"OK", "wordlist/testdata/testdata.txt", []string{"test1", "test2", "test3"}, nil},
		{"NG: cannot find text", "wordlist/testdata/noexists.txt", nil, xerrors.New("open wordlist/testdata/noexists.txt: no such file or directory")},
	}

	for _, test := range runImportTests {

		t.Run(test.testCase, func(t *testing.T) {
			actual, err := Import(test.path)
			if !reflect.DeepEqual(err, test.err) {
				if err.Error() != test.err.Error() { // この部分をどうすべきか
					t.Fatalf("Failed test. expected: %s, but actual: %s", test.err, err)
				}
			}

			if !reflect.DeepEqual(actual, test.out) {
				t.Fatalf("Failed test. expected: %s, but actual: %s", test.out, actual)
			}

		})
	}
}
