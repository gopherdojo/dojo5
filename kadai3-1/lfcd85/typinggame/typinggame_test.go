package typinggame_test

import (
	"bytes"
	"regexp"
	"testing"
	"time"

	"github.com/gopherdojo/dojo5/kadai3-1/lfcd85/typinggame"
)

func TestExecute(t *testing.T) {
	g := typinggame.Game{
		typinggame.Words{"hoge"},
		1 * time.Second,
	}

	if err := typinggame.Execute(g); err != nil {
		t.Errorf("failed to execute new game: %v", err)
	}
}

func TestGame_run(t *testing.T) {
	ch := make(chan string)
	go func() {
		time.Sleep(100 * time.Millisecond)
		ch <- "hoga"
		time.Sleep(100 * time.Millisecond)
		ch <- "hoge"
	}()

	g := typinggame.Game{
		typinggame.Words{"hoge"},
		1 * time.Second,
	}

	var output bytes.Buffer
	typinggame.ExportGameRun(&g, ch, &output)

	cases := []struct {
		output   string
		expected bool
	}{
		{"hoga ... NG", true},
		{"hoga ... OK", false},
		{"hoge ... OK", true},
		{"hoge ... NG", false},
	}

	for _, c := range cases {
		c := c
		t.Run(c.output, func(t *testing.T) {
			t.Parallel()

			actual := regexp.MustCompile(c.output).MatchString(output.String())
			if actual != c.expected {
				switch c.expected {
				case true:
					t.Errorf("%v should be outputted but actually was not", c.output)
				case false:
					t.Errorf("%v should not be outputted but actyally was", c.output)
				}
			}
		})
	}
}
