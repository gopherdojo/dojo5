package typinggame_test

import (
	"testing"
	"time"

	"github.com/gopherdojo/dojo5/kadai3-1/lfcd85/typinggame"
)

func TestExecute(t *testing.T) {
	g := typinggame.Game{
		typinggame.Words{"hoge", "fuga", "piyo"},
		30 * time.Second,
	}

	if err := typinggame.Execute(g); err != nil {
		t.Errorf("failed to execute new game: %v", err)
	}
}
