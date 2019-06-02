package walk

import (
	"log"
	"os"
	"reflect"
	"testing"
)

func TestFind(t *testing.T) {

	// setup
	//
	// test
	// |- gopher.png
	// |- gopher.jpg
	// |- gopher.gif
	// |- subdir
	//    |- gopher-sub.png
	//    |- gopher-sub.jpg
	//    |- gopher.jpg.png
	//    |- gopher.png.jpg
	//    |- gopher-symlink.png -> ../gopher.gif
	err := os.MkdirAll("./test/subdir", 0777)
	if err != nil {
		log.Fatal("can't make test dir")
	}
	_, err = os.Create("./test/gopher.png")
	if err != nil {
		log.Fatal("can't make test data")
	}
	_, err = os.Create("./test/gopher.jpg")
	if err != nil {
		log.Fatal("can't make test data")
	}
	_, err = os.Create("./test/gopher.gif")
	if err != nil {
		log.Fatal("can't make test data")
	}
	_, err = os.Create("./test/subdir/gopher-sub.png")
	if err != nil {
		log.Fatal("can't make test data")
	}
	_, err = os.Create("./test/subdir/gopher-sub.jpg")
	if err != nil {
		log.Fatal("can't make test data")
	}
	_, err = os.Create("./test/subdir/gopher.jpg.png")
	if err != nil {
		log.Fatal("can't make test data")
	}
	_, err = os.Create("./test/subdir/gopher.png.jpg")
	if err != nil {
		log.Fatal("can't make test data")
	}
	err = os.Symlink("./test/gopher.gif", "./test/subdir/gopher-symlink.png")
	if err != nil {
		log.Fatal("can't make test data")
	}

	// main test
	w := NewWalker()
	paths, err := w.Find("./test/", "jpg")
	if err != nil {
		t.Errorf("Failed : %s", err)
	}
	check := []string{"test/gopher.jpg", "test/subdir/gopher-sub.jpg", "test/subdir/gopher.png.jpg"}
	if !reflect.DeepEqual(paths, check) { // reflectしてるから遅そう…
		t.Errorf("Failed: expect %s but actual %s", check, paths)
	}

	paths, err = w.Find("./test/subdir", "png")
	if err != nil {
		t.Errorf("Failed : %s", err)
	}
	check = []string{"test/subdir/gopher-sub.png", "test/subdir/gopher.jpg.png"}
	if !reflect.DeepEqual(paths, check) { // reflectしてるから遅そう…
		t.Errorf("Failed: expect %s but actual %s", check, paths)
	}

	paths, err = w.Find("", "jpg")
	if err == nil {
		t.Errorf("Failed : expect err is non nil but actual err is nil")
	}
	if paths != nil {
		t.Errorf("Failed: expect paths is nil but %s", paths)
	}

	paths, err = w.Find("test", "")
	if err != nil {
		t.Errorf("Failed : %s", err)
	}
	check = make([]string, 0)
	if !reflect.DeepEqual(paths, check) { // reflectしてるから遅そう…
		t.Errorf("Failed: expect %s but actual %s", check, paths)
	}

	// clean
	err = os.RemoveAll("./test/")
	if err != nil {
		log.Fatal("can't clean test data")
	}
}
