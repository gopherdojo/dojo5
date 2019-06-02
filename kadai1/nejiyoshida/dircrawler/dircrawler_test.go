package dircrawler

import (
	//"fmt"

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
	testCase1 := []string{"a.txt",
		"b.jpg",
		"c.jpg",
		"d.html"}
	ret1 := []string{"b.jpg", "c.jpg"}
	if !reflect.DeepEqual(selectFormat(testCase1, ".jpg"), ret1) {
		t.Fatal("testCase1 does not match")
	}

}

func TestIsSpecifiedFormat(t *testing.T) {

	if !isSpecifiedFormat("./hoge/huga.jpg", ".jpg") {
		t.Fatal("isSpecifiedFormat(\"./hoge/huga.jpg\" , \"jpg\") should be true, but false")
	}
}
func TestSearchFilePaths(t *testing.T) {
	ret := []string{"test\\TEST\\bar.txt", "test\\TEST\\gopher.jpg", "test\\foo.txt"}
	if !reflect.DeepEqual(searchFilePaths("./test"), ret) {
		t.Fatal("does not match")

	}
}
