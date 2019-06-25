package typinggame_test

import (
	"testing"

	"github.com/gopherdojo/dojo5/kadai3-1/manhdaovan/pkg/typinggame"
)

func Test_defaultSentencePickFnc(t *testing.T) {
	type args struct {
		sentences []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"nil given",
			args{nil},
			"",
		},
		{
			"no sentence given",
			args{[]string{}},
			"",
		},
		{
			"just one sentence given",
			args{[]string{"sample sentence"}},
			"sample sentence",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := typinggame.DefaultSentencePickFnc(tt.args.sentences); got != tt.want {
				t.Errorf("defaultSentencePickFnc() = %v, want %v", got, tt.want)
			}
		})
	}
}
