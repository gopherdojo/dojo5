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
var defaultSentences = []string{
	"Joe made the sugar cookies; Susan decorated them.",
	"There were white out conditions in the town; subsequently, the roads were impassable.",
	"She folded her handkerchief neatly.",
	"Two seats were vacant.",
	"She advised him to come back at once.",
	"Sometimes, all you need to do is completely make an ass of yourself and laugh it off to realise that life isn’t so bad after all.",
	"We have a lot of rain in June.",
	"Rock music approaches at high velocity.",
	"Everyone was busy, so I went to the movie alone.",
	"We have never been to Asia, nor have we visited Africa.",
	"Check back tomorrow; I will see if the book has arrived.",
	"Wednesday is hump day, but has anyone asked the camel if he’s happy about it?",
	"The book is in front of the table.",
	"Let me help you with your baggage.",
	"Please wait outside of the house.",
	"She wrote him a long letter, but he didn't read it.",
	"I want to buy a onesie… but know it won’t suit me.",
	"Lets all be unique together until we realise we are all the same.",
	"He told us a very exciting adventure story.",
	"Yeah, I think it's a good environment for learning English.",
}

type PickerFnc func(sentences []string) string

type TypingGame struct {
	Duration     uint64 // seconds
	Sentences    []string
	PickSentence PickerFnc
	QuitSigs     []os.Signal
	doneChan     chan string
	sigChan      chan os.Signal
	errChan      chan error
	textIn       io.Reader
	textOut      io.Writer
	correctNum   int
}

func (tg *TypingGame) Start() <-chan string {
	tg.start()
	return tg.doneChan
}

func (tg *TypingGame) CorrectSentences() int {
	return tg.correctNum
}

func (tg *TypingGame) start() {
	tg.initGame()
	tg.listenQuitSig()

	go tg.play()

	ctx, _ := context.WithTimeout(context.Background(), time.Duration(tg.Duration)*time.Second)
	select {
	case <-ctx.Done():
		tg.doneChan <- "Time is up!"
	case sig := <-tg.sigChan:
		tg.doneChan <- fmt.Sprintf("Got quit sig: %s!", sig.String())
	case err := <-tg.errChan:
		tg.doneChan <- fmt.Sprintf("Got error: %v", err)
	}

	close(tg.doneChan)
	close(tg.sigChan)
	close(tg.errChan)
}

func (tg *TypingGame) initGame() {
	tg.errChan = make(chan error, 1)
	tg.doneChan = make(chan string, 1)
	tg.sigChan = make(chan os.Signal, len(tg.QuitSigs)+len(defaultSigsQuit))

	tg.textIn = os.Stdin
	tg.textOut = os.Stdout

	if tg.PickSentence == nil {
		tg.PickSentence = defaultSentencePickFnc
	}

	if len(tg.Sentences) == 0 {
		tg.Sentences = defaultSentences
	}
}

func (tg *TypingGame) play() {
	scanner := bufio.NewScanner(tg.textIn)
	for {
		sampleSentence := tg.PickSentence(tg.Sentences)
		if err := tg.print([]byte(sampleSentence + "\n")); err != nil {
			tg.errChan <- err
			return
		}

		if !scanner.Scan() {
			tg.errChan <- scanner.Err()
			return
		}
		inputSentence := scanner.Bytes()
		if string(inputSentence) == sampleSentence {
			tg.correctNum++
		}

		select {
		case <-tg.doneChan:
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
