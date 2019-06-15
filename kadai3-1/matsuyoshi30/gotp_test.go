package main

import "testing"

func Test_gotp(t *testing.T) {
	if err := gotp(1, "default.txt"); err != nil {
		t.Fatalf("ERROR: %v", err)
	}
}
