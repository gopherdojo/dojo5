package typinggame

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"syscall"
	"testing"
	"time"
)

func TestTypingGame_Start(t *testing.T) {
	fireTimeoutFnc := func(timeC chan struct{}) {
		timeC <- struct{}{}
	}
	fireSigFnc := func(sigC chan os.Signal) {
		sigC <- syscall.SIGINT
	}
	fireErrFnc := func(errC chan error) {
		errC <- fmt.Errorf("error while read/write to io")
	}

	tests := []struct {
		name           string
		fireTimeoutFnc func(timeC chan struct{})
		fireSigFnc     func(sigC chan os.Signal)
		fireErrFnc     func(errC chan error)
		wantExitReason string
	}{
		{
			name:           "exit by timeout",
			fireTimeoutFnc: fireTimeoutFnc,
			fireSigFnc:     nil,
			fireErrFnc:     nil,
			wantExitReason: "Time is up!",
		},
		{
			name:           "exit by signal",
			fireTimeoutFnc: nil,
			fireSigFnc:     fireSigFnc,
			fireErrFnc:     nil,
			wantExitReason: "Got quit sig: interrupt!",
		},
		{
			name:           "exit by error",
			fireTimeoutFnc: nil,
			fireSigFnc:     nil,
			fireErrFnc:     fireErrFnc,
			wantExitReason: "Got error: error while read/write to io",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			timeC := make(chan struct{}, 1)
			sigC := make(chan os.Signal, 2)
			errC := make(chan error, 3)
			tg := &TypingGame{
				Duration:     1 * time.Second,
				PickSentence: func([]string) string { return "" },
				textIn:       os.Stdin,
				textOut:      os.Stdout,
				timeChan:     timeC,
				sigChan:      sigC,
				errChan:      errC,
			}

			go func() {
				if tt.fireTimeoutFnc != nil {
					tt.fireTimeoutFnc(timeC)
				}
				if tt.fireSigFnc != nil {
					tt.fireSigFnc(sigC)
				}
				if tt.fireErrFnc != nil {
					tt.fireErrFnc(errC)
				}
			}()

			reason := tg.waitExit()
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
	initPicker := func() PickerFnc {
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
				errChan: make(chan error, 2),
				textIn:  &mockReader{corrected: 0},
				textOut: &bytes.Buffer{},
			},
			want: 0,
		},
		{
			name: "1 corrected sentence",
			fields: fields{
				errChan: make(chan error, 2),
				textIn:  &mockReader{corrected: 1},
				textOut: &bytes.Buffer{},
			},
			want: 1,
		},
		{
			name: "2 corrected sentences",
			fields: fields{
				errChan: make(chan error, 2),
				textIn:  &mockReader{corrected: 2},
				textOut: &bytes.Buffer{},
			},
			want: 2,
		},
		{
			name: "error on io.Reader",
			fields: fields{
				errChan: make(chan error, 2),
				textIn:  &mockReader{err: fmt.Errorf("error on io.Reader")},
				textOut: &bytes.Buffer{},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tg := &TypingGame{
				Sentences:    sampleSentence,
				PickSentence: initPicker(),
				errChan:      tt.fields.errChan,
				textIn:       tt.fields.textIn,
				textOut:      tt.fields.textOut,
			}

			tg.play(context.Background())
			if got := tg.CorrectSentences(); got != tt.want {
				t.Errorf("TypingGame.CorrectSentences() = %v, want %v", got, tt.want)
			}
		})
	}
}
