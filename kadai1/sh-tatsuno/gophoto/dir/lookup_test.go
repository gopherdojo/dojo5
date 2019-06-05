package dir

import (
	"reflect"
	"testing"
)

func Test_Lookup(t *testing.T) {

	t.Run("OK: .jpg", func(t *testing.T) {
		var actual []string
		// ### Given ###
		actual, err := Lookup("./test", ".jpg", actual)
		if err != nil {
			t.Fatalf("Failed test. err: %v", err)
		}

		expected := []string{"test/test1.jpg", "test/test3/test4.jpg", "test/test3/test5.jpg"}

		// ### Then ###
		if !reflect.DeepEqual(actual, expected) {
			t.Fatalf("Failed test. expected: %v,\n but actual: %v", expected, actual)
		}
	})

	t.Run("OK: .png", func(t *testing.T) {
		var actual []string
		// ### Given ###
		actual, err := Lookup("./test", ".png", actual)
		if err != nil {
			t.Fatalf("Failed test. err: %v", err)
		}

		expected := []string{"test/test2.png", "test/test3/test6.png"}

		// ### Then ###
		if !reflect.DeepEqual(actual, expected) {
			t.Fatalf("Failed test. expected: %v,\n but actual: %v", expected, actual)
		}
	})

	t.Run("OK: .gif(empty)", func(t *testing.T) {
		var actual []string
		// ### Given ###
		actual, err := Lookup("./test", ".gif", actual)
		if err != nil {
			t.Fatalf("Failed test. err: %v", err)
		}

		var expected []string

		// ### Then ###
		if !reflect.DeepEqual(actual, expected) {
			t.Fatalf("Failed test. expected: %v,\n but actual: %v", expected, actual)
		}
	})

	t.Run("NG: not found directory", func(t *testing.T) {
		var actual []string
		// ### Given ###
		_, err := Lookup("./test2", ".jpg", actual)
		if err == nil {
			t.Fatalf("should be error.")
		}

		expected := "open ./test2: no such file or directory"

		// ### Then ###
		if err.Error() != expected {
			t.Fatalf("Failed test. expected: %s,\n but actual: %s", expected, err.Error())
		}
	})
}
