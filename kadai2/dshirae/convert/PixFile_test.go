package convert

import "testing"

// picture file 変換テスト
func TestPixFile(t *testing.T) {

	testTables := []PixConv{
		{"./test.jpg", "jpeg", "png"},
		{"./test.jpg", "jpeg", "gif"},
		{"./test.png", "png", "jpeg"},
		{"./test.png", "png", "gif"},
		{"./test.gif", "gif", "jpeg"},
		{"./test.gif", "gif", "png"},
	}

	for _, testpix := range testTables {
		testPixFile(t, testpix)
	}
}

func testPixFile(t *testing.T, testpix PixConv) {
	t.Helper()
	err := PixFile(testpix)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
}
