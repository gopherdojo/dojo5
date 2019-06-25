package typinggame

import (
	"context"
	"io"
	"os"
	"time"
)

// DefaultSentencePickFnc exports defaultSentencePickFnc for testing
var DefaultSentencePickFnc = defaultSentencePickFnc

// NewTypingGameForTest init new TypingGame struct for testing
func NewTypingGameForTest(ctx context.Context, sentences []string,
	pickSentence PickerFnc, quitSigs []os.Signal,
	sigChan chan os.Signal, errChan chan error, duration time.Duration,
	textIn io.Reader, textOut io.Writer) *TypingGame {

	return &TypingGame{
		Duration:     duration,
		Sentences:    sentences,
		PickSentence: pickSentence,
		QuitSigs:     quitSigs,
		sigChan:      sigChan,
		errChan:      errChan,
		ctx:          ctx,
		textIn:       textIn,
		textOut:      textOut,
	}
}

// WaitExitForTest exports waitExit for testing
func (tg *TypingGame) WaitExitForTest() string {
	return tg.waitExit()
}

// PlayForTest exports play method for testing
func (tg *TypingGame) PlayForTest(ctx context.Context) {
	tg.play(ctx)
}
