package main

import (
	"bytes"
	"testing"
)

func TestJudge(t *testing.T) {
	tests := []struct {
		name  string
		q     string
		a     string
		score int
	}{
		{
			name:  "case1",
			q:     "abc",
			a:     "abc",
			score: 3,
		},
		{
			name:  "case2",
			q:     "abc",
			a:     "ab",
			score: 2,
		},
		{
			name:  "case3",
			q:     "ab",
			a:     "abc",
			score: 2,
		},
		{
			name:  "case4",
			q:     "",
			a:     "abc",
			score: 0,
		},
		{
			name:  "case5",
			q:     "abc",
			a:     "",
			score: 0,
		},
		{
			name:  "case6",
			q:     "",
			a:     "",
			score: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := judge(test.q, test.a); got != test.score {
				t.Errorf("got %v, want %v", got, test.score)
			}
		})
	}
}

func TestInput(t *testing.T) {
	tests := []struct {
		name string
		in   string
		out  string
	}{
		{
			name: "case1",
			in:   "abc",
			out:  "abc",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			r := bytes.NewBuffer([]byte(test.in))
			if got := <-input(r); got != test.out {
				t.Errorf("got %v, want %v", got, test.out)
			}
		})
	}
}
