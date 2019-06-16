package dir

import (
	"reflect"
	"testing"
)

func Test_Lookup(t *testing.T) {
	t.Helper()

	lookupOKTests := []struct {
		testCase string
		ext      string
		expected []string
	}{
		{"OK: .jpg", ".jpg", []string{"testdata/test1.jpg", "testdata/test3/test4.jpg", "testdata/test3/test5.jpg"}},
		{"OK: .png", ".png", []string{"testdata/test2.png", "testdata/test3/test6.png"}},
		{"OK: .gif", ".gif", []string{}},
	}
	for _, test := range lookupOKTests {

		t.Run(test.testCase, func(t *testing.T) {
			actual := []string{}
			// ### Given ###
			actual, err := Lookup("./testdata", test.ext, actual)
			if err != nil {
				t.Fatalf("Failed test. err: %v", err)
			}

			// ### Then ###
			if !reflect.DeepEqual(actual, test.expected) {
				t.Fatalf("Failed test. expected: %v,\n but actual: %v", test.expected, actual)
			}
		})
	}

	t.Run("NG: not found directory", func(t *testing.T) {
		actual := []string{}
		// ### Given ###
		_, err := Lookup("./testdata2", ".jpg", actual)
		if err == nil {
			t.Fatalf("should be error.")
		}

		expected := "open ./testdata2: no such file or directory"

		// ### Then ###
		if err.Error() != expected {
			t.Fatalf("Failed test. expected: %s,\n but actual: %s", expected, err.Error())
		}
	})
}
