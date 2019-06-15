package questions_test

import (
	"fmt"
	"log"
	"math/rand"
	"reflect"
	"testing"
	"time"

	"github.com/gopherdojo/dojo5/kadai3-1/nagaa052/pkg/questions"
	qs "github.com/gopherdojo/dojo5/kadai3-1/nagaa052/pkg/questions"
)

func Example() {
	qs, err := questions.New()
	if err != nil {
		log.Fatal("Failed new quetions")
	}

	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(qs.GetSize())

	q, err := qs.GetOne(index)
	if err != nil {
		log.Fatal("Failed get quetion")
	}

	answer := q.Word
	if q.IsCorrect(answer) {
		fmt.Println("Correct!")
	}
	// Output:
	// Correct!
}

func TestNew(t *testing.T) {

	imq, err := qs.ExportNewInMemQ()
	if err != nil {
		t.Errorf("Failed new In memoruy quetions")
	}

	tests := []struct {
		name    string
		want    *qs.Questions
		wantErr bool
	}{
		{
			name: "Success test",
			want: &qs.Questions{
				Store: imq,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := qs.New()
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewWithStore(t *testing.T) {
	type args struct {
		s qs.Store
	}
	tests := []struct {
		name    string
		args    args
		want    *qs.Questions
		wantErr bool
	}{
		{
			name: "Success test",
			args: args{},
			want: &qs.Questions{
				Store: args{}.s,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := qs.NewWithStore(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewWithStore() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewWithStore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuestion_IsCorrect(t *testing.T) {
	type fields struct {
		Word string
	}
	tests := []struct {
		name   string
		fields fields
		answer string
		want   bool
	}{
		{
			name: "Is correct test",
			fields: fields{
				Word: "HogeHoge",
			},
			answer: "HogeHoge",
			want:   true,
		},
		{
			name: "Is not correct test",
			fields: fields{
				Word: "FugaFuga",
			},
			answer: "fugafuga",
			want:   false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			q := &qs.Question{
				Word: tt.fields.Word,
			}
			if got := q.IsCorrect(tt.answer); got != tt.want {
				t.Errorf("Question.IsCorrect() = %v, want %v", got, tt.want)
			}
		})
	}
}
