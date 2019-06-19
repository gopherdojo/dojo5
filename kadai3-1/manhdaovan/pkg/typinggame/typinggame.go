package typinggame

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var defaultSigsQuit = []os.Signal{syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT}

// PickerFnc is method to pick a sentence from sample sentences
type PickerFnc func(sentences []string) string

// TypingGame represents the game struct
type TypingGame struct {
	Duration     time.Duration
	Sentences    []string
	PickSentence PickerFnc
	QuitSigs     []os.Signal

	sigChan      chan os.Signal
	errChan      chan error
	timeChan     <- chan struct{}

	textIn       io.Reader
	textOut      io.Writer

	correctNum   int
}

// Start starts the game
func (tg *TypingGame) Start() string {
	ctx, cancel := context.WithTimeout(context.Background(), tg.Duration)
	defer cancel()

	if err := tg.initGame(ctx.Done()); err != nil {
		return fmt.Sprintf("Error on init game: %v", err)
	}

	tg.listenQuitSig()
	go tg.play(ctx)
	return tg.waitExit()
}

// CorrectSentences returns number of corrected sentences from input
func (tg *TypingGame) CorrectSentences() int {
	return tg.correctNum
}

func (tg *TypingGame) waitExit() string {
	var exitReason string
	select {
	case <-tg.timeChan:
		exitReason = "Time is up!"
	case sig := <-tg.sigChan:
		exitReason = fmt.Sprintf("Got quit sig: %s!", sig.String())
	case err := <-tg.errChan:
		exitReason = fmt.Sprintf("Got error: %v", err)
	}

	close(tg.sigChan)
	close(tg.errChan)

	return exitReason
}

func (tg *TypingGame) initGame(c <- chan struct{}) error {
	tg.timeChan = c
	tg.errChan = make(chan error, 2)     // cap for 1 error and closing channel data
	tg.sigChan = make(chan os.Signal, 2) // cap for 1 signal and closing channel data

	tg.textIn = os.Stdin
	tg.textOut = os.Stdout

	if tg.PickSentence == nil {
		tg.PickSentence = defaultSentencePickFnc
	}

	if len(tg.Sentences) == 0 { // no given sample sentences
		return fmt.Errorf("no sample sentences given")
	}

	return nil
}

func (tg *TypingGame) play(ctx context.Context) {
	scanner := bufio.NewScanner(tg.textIn)

	for {
		sampleSentence := tg.PickSentence(tg.Sentences)
		if err := tg.print([]byte(sampleSentence + "\n")); err != nil {
			tg.errChan <- err
			return
		}

		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				tg.errChan <- scanner.Err()
			}
			return
		}

		inputSentence := scanner.Bytes()
		if string(inputSentence) == sampleSentence {
			tg.correctNum++
		}

		select {
		case <-ctx.Done():
			return
		default:
		}
	}
}

func (tg *TypingGame) print(sentence []byte) error {
	_, err := tg.textOut.Write(sentence)
	return err
}

func (tg *TypingGame) listenQuitSig() {
	// listen to default signals
	tg.QuitSigs = append(tg.QuitSigs, defaultSigsQuit...)
	signal.Notify(tg.sigChan, tg.QuitSigs...)
}
