package typinggame

import "testing"

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
			if got := defaultSentencePickFnc(tt.args.sentences); got != tt.want {
				t.Errorf("defaultSentencePickFnc() = %v, want %v", got, tt.want)
			}
		})
	}
}
