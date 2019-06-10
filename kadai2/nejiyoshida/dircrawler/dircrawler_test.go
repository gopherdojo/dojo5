package dircrawler

import (
	"fmt"
	"reflect"
	"testing"
)

func ExampleisSpecifiedFormat() {

	fmt.Println(isSpecifiedFormat("hoge.jpg", ".jpg"))
	//Output:true
	fmt.Println(isSpecifiedFormat("huga.jpg", ".png"))
	//Output:false
	fmt.Println(isSpecifiedFormat("foo.png", ".png"))
	//Output:true
	fmt.Println(isSpecifiedFormat("bar.png", ".png"))
	//Output:false
}

func TestSelectFormat(t *testing.T) {
	t.Helper()
	cases := []struct {
		ext   string
		files []string
		ret   []string
	}{
		{ext: ".jpg", files: []string{"a.txt", "b.jpg", "c.jpg", "d.html", "e.png"}, ret: []string{"b.jpg", "c.jpg"}},
		{ext: ".exe", files: []string{"a.txt", "b.jpg", "c.jpg", "d.html", "e.png"}, ret: nil},
		{ext: ".jpg", files: []string{"./hoge/a.txt", "./hoge/huga/b.jpg", "./.git/c.jpg", "./d.html", "./e.png"}, ret: []string{"./hoge/huga/b.jpg", "./.git/c.jpg"}},
	}
	for _, c := range cases {
		if !reflect.DeepEqual(selectFormat(c.files, c.ext), c.ret) {
			t.Errorf("test cases %v does not match\n", c)
		}
	}

}

func TestIsSpecifiedFormat(t *testing.T) {
	t.Helper()
	cases := []struct {
		file string
		ext  string
		ans  bool
	}{
		{file: "a.jpg", ext: ".jpg", ans: true},
		{file: "a.jpg", ext: ".txt", ans: false},
		{file: "./hoge/huga/a.png", ext: ".png", ans: true},
		{file: "./foo/bar/a.jpg", ext: ".png", ans: false},
	}
	for _, c := range cases {
		if !isSpecifiedFormat("./hoge/huga.jpg", ".jpg") {
			t.Errorf("test cases %v is not %v formt\n", c.file, c.ext)
		}
	}
}
func TestSearchFilePaths(t *testing.T) {
	t.Helper()
	cases := []struct {
		rootDir string
		ret     []string
	}{
		{rootDir: "./test", ret: []string{"test\\TEST\\bar.txt", "test\\TEST\\gopher.jpg", "test\\foo.txt"}},
		{rootDir: "./test/TEST", ret: []string{"test\\TEST\\bar.txt", "test\\TEST\\gopher.jpg"}},
	}
	for _, c := range cases {

		if !reflect.DeepEqual(searchFilePaths(c.rootDir), c.ret) {
			t.Errorf("test case %v does not match\n", c)
		}

	}
}
