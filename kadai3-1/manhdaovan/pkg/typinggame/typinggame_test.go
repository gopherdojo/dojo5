package typinggame

import (
	"fmt"
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

	type fields struct {
		doneChan chan string
		sigChan  chan os.Signal
		errChan  chan error
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
