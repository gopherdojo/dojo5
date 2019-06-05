package walker

import (
	"fmt"
	"testing"
)

func Test_Walk(t *testing.T) {
	ch, err := Walk("./testdir")
	if err != nil {
		t.Fatal(err)
	}

	count := 0
	for filePath := range ch {
		fmt.Println(filePath)
		count++
	}

	if count != 3 {
		fmt.Println(len(ch))
		t.Fatal("failed test")
	}
}
