package typinggame_test

import (
	"testing"

	"github.com/gopherdojo/dojo5/kadai3-1/lfcd85/typinggame"
)

func TestExecute(t *testing.T) {
	if err := typinggame.Execute(); err != nil {
		t.Errorf("error: %v", err)
	}
}
