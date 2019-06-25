package typinggame_test

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"sync"
	"syscall"
	"testing"
	"time"

	"github.com/gopherdojo/dojo5/kadai3-1/manhdaovan/pkg/typinggame"
)

func TestTypingGame_Start(t *testing.T) {
	fireSigFnc := func(sigC chan os.Signal) {
		sigC <- syscall.SIGINT
	}
	fireErrFnc := func(errC chan error) {
		errC <- fmt.Errorf("error while read/write to io")
	}

	tests := []struct {
		name           string
		fireTimeout    bool
		fireSigFnc     func(sigC chan os.Signal)
		fireErrFnc     func(errC chan error)
		wantExitReason string
	}{
		{
			name:           "exit by timeout",
			fireTimeout:    true,
			fireSigFnc:     nil,
			fireErrFnc:     nil,
			wantExitReason: "Time is up!",
		},
		{
			name:           "exit by signal",
			fireTimeout:    false,
			fireSigFnc:     fireSigFnc,
			fireErrFnc:     nil,
			wantExitReason: "Got quit sig: interrupt!",
		},
		{
			name:           "exit by error",
			fireTimeout:    false,
			fireSigFnc:     nil,
			fireErrFnc:     fireErrFnc,
			wantExitReason: "Got error: error while read/write to io",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			sigC := make(chan os.Signal, 2)
			errC := make(chan error, 3)
			tg := typinggame.NewTypingGameForTest(ctx, nil,
				func([]string) string { return "" }, nil,
				sigC, errC, 1*time.Second,
				os.Stdin, os.Stdout)

			var wg sync.WaitGroup
			wg.Add(1) // prevent race condition
			go func() {
				if tt.fireTimeout {
					cancel()
				}
				if tt.fireSigFnc != nil {
					tt.fireSigFnc(sigC)
				}
				if tt.fireErrFnc != nil {
					tt.fireErrFnc(errC)
				}
				wg.Done()
			}()
			wg.Wait()

			reason := tg.WaitExitForTest()
			if reason != tt.wantExitReason {
				t.Errorf("TypingGame.Start() = %v, want %v", reason, tt.wantExitReason)
			}
		})
	}
}

var sampleSentence = []string{
	"sample sentence 1",
	"sample sentence 2",
	"sample sentence 3",
}

type mockReader struct {
	corrected int
	readTurn  int
	err       error
}

func (m *mockReader) Read(p []byte) (int, error) {
	if m.err != nil {
		return 0, m.err
	}
	if m.readTurn == m.corrected {
		return 0, nil // read done
	}

	// read next sentence from sampleSentence
	sentence := []byte(sampleSentence[m.readTurn])
	for i, b := range sentence {
		p[i] = b
	}
	p[len(sentence)] = '\n' // simulate newline character

	m.readTurn++

	return len(sentence) + 1, nil
}

func TestTypingGame_CorrectSentences(t *testing.T) {
	initPicker := func() typinggame.PickerFnc {
		idx := -1

		return func(sentences []string) string {
			idx++
			if len(sentences) == 0 || idx >= len(sentences) {
				return ""
			}
			return sentences[idx]
		}
	}

	type fields struct {
		errChan chan error
		textIn  io.Reader
		textOut io.Writer
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "no corrected sentence",
			fields: fields{
				errChan: make(chan error, 1),
				textIn:  &mockReader{corrected: 0},
				textOut: &bytes.Buffer{},
			},
			want: 0,
		},
		{
			name: "1 corrected sentence",
			fields: fields{
				errChan: make(chan error, 1),
				textIn:  &mockReader{corrected: 1},
				textOut: &bytes.Buffer{},
			},
			want: 1,
		},
		{
			name: "2 corrected sentences",
			fields: fields{
				errChan: make(chan error, 1),
				textIn:  &mockReader{corrected: 2},
				textOut: &bytes.Buffer{},
			},
			want: 2,
		},
		{
			name: "error on io.Reader",
			fields: fields{
				errChan: make(chan error, 1),
				textIn:  &mockReader{err: fmt.Errorf("error on io.Reader")},
				textOut: &bytes.Buffer{},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			tg := typinggame.NewTypingGameForTest(ctx,
				sampleSentence,
				initPicker(),
				nil, nil,
				tt.fields.errChan, 0,
				tt.fields.textIn,
				tt.fields.textOut)

			tg.PlayForTest(ctx)
			if got := tg.CorrectSentences(); got != tt.want {
				t.Errorf("TypingGame.CorrectSentences() = %v, want %v", got, tt.want)
			}
		})
	}
}
