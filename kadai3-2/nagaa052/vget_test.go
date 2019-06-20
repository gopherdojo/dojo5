package vget_test

import (
	"bytes"
	"testing"

	vget "github.com/gopherdojo/dojo5/kadai3-2/nagaa052"
)

func TestNew(t *testing.T) {
	type args struct {
		url string
		opt vget.Options
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Success Test",
			args: args{
				url: "http://example.com",
				opt: vget.DefaultOptions,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			outStream := &bytes.Buffer{}
			errStream := &bytes.Buffer{}
			_, err := vget.New(tt.args.url, tt.args.opt, outStream, errStream)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
