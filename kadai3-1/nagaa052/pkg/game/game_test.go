package game_test

import (
	"bytes"
	"io"
	"log"
	"reflect"
	"testing"

	"github.com/gopherdojo/dojo5/kadai3-1/nagaa052/pkg/game"
	"github.com/gopherdojo/dojo5/kadai3-1/nagaa052/pkg/questions"
	qs "github.com/gopherdojo/dojo5/kadai3-1/nagaa052/pkg/questions"
)

func TestResult_Print(t *testing.T) {
	successOut := `========================
Correct Count: 1
Correct Rate: 50.0％
`

	type fields struct {
		Questions    []*qs.Question
		CorrectCount int
	}
	tests := []struct {
		name    string
		fields  fields
		wantOut string
	}{
		{
			name: "Success test",
			fields: fields{
				Questions:    []*qs.Question{&qs.Question{}, &qs.Question{}},
				CorrectCount: 1,
			},
			wantOut: successOut,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &game.Result{
				Questions:    tt.fields.Questions,
				CorrectCount: tt.fields.CorrectCount,
			}
			out := &bytes.Buffer{}
			r.Print(out)
			if gotOut := out.String(); gotOut != tt.wantOut {
				t.Errorf("Result.Print() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		opt      game.Options
		inStream io.Reader
	}
	tests := []struct {
		name    string
		opt     game.Options
		wantErr bool
	}{
		{
			name:    "Success Test",
			opt:     game.DefaultOptions,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			inStream := &bytes.Buffer{}
			outStream := &bytes.Buffer{}
			errStream := &bytes.Buffer{}

			_, err := game.New(tt.opt, inStream, outStream, errStream)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestGame_getQuestion(t *testing.T) {
	tests := []struct {
		name     string
		opt      game.Options
		wantWord string
		wantErr  bool
	}{
		{
			name:     "Success Test",
			opt:      game.DefaultOptions,
			wantWord: "Hoge",
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := testNewMock(t, tt.opt)
			got, err := game.ExportGetQuestion(g)
			if (err != nil) != tt.wantErr {
				t.Errorf("Game.getQuestion() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.Word, tt.wantWord) {
				t.Errorf("Game.getQuestion() = %v, want %v", got, tt.wantWord)
			}
		})
	}
}

type MockStore struct{}

func (ms *MockStore) GetSize() int {
	return 1
}

func (ms *MockStore) GetOne(index int) (*questions.Question, error) {
	return &questions.Question{
		Word: "Hoge",
	}, nil
}

func testNewMock(t *testing.T, opt game.Options) *game.Game {
	t.Helper()

	inStream := &bytes.Buffer{}
	outStream := &bytes.Buffer{}
	errStream := &bytes.Buffer{}

	g, err := game.New(opt, inStream, outStream, errStream)
	if err != nil {
		log.Fatal("Failed to generate game mock")
	}

	mockStore, err := questions.NewWithStore(&MockStore{})
	if err != nil {
		log.Fatal("Failed to generate store mock")
	}

	g.ExportSetQs(mockStore)
	return g
}
