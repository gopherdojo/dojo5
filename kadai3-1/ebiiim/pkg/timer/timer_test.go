package timer_test

import (
	"testing"

	"github.com/gopherdojo/dojo5/kadai3-1/ebiiim/pkg/timer"
)

func TestMakeChannel(t *testing.T) {
	cases := []struct {
		name     string
		timeSec  int
		isClosed bool
	}{
		{name: "normal", timeSec: 2, isClosed: true},
		{name: "+boundary", timeSec: 1, isClosed: true},
		{name: "0sec", timeSec: 0, isClosed: true},
		{name: "-boundary", timeSec: -1, isClosed: true},
		{name: "-N", timeSec: -2, isClosed: true},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			timerCh := timer.MakeChannel(c.timeSec)
			_, isNotClosed := <-timerCh
			if isNotClosed == c.isClosed {
				t.Error("channel status error")
			}
		})
	}
}
