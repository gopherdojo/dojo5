package questions_test

import (
	"reflect"
	"testing"

	q "github.com/gopherdojo/dojo5/kadai3-1/nagaa052/pkg/questions"
)

func Test_inMemQ_GetOne(t *testing.T) {
	imq, err := q.ExportNewInMemQ()
	if err != nil {
		t.Errorf("Failed new In memoruy quetions")
	}

	type fields struct {
		words []string
	}
	type args struct {
		index int
	}
	tests := []struct {
		name    string
		args    args
		want    *q.Question
		wantErr bool
	}{
		{
			name: "Success test",
			args: args{
				index: 0,
			},
			want: &q.Question{
				Word: q.ExportInMemWords[0],
			},
			wantErr: false,
		},
		{
			name: "Question not found test",
			args: args{
				index: 9999999,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := imq.GetOne(tt.args.index)
			if (err != nil) != tt.wantErr {
				t.Errorf("inMemQ.GetOne() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inMemQ.GetOne() = %v, want %v", got, tt.want)
			}
		})
	}
}
