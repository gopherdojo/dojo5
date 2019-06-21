package timer

import (
	"time"
)

// MakeChannel returns a read-only channel that closes after given seconds.
func MakeChannel(sec int) <-chan interface{} {
	timer := make(chan interface{})
	go func() {
		//defer fmt.Println("closed TimerChannel")
		defer close(timer)
		<-time.After(time.Duration(sec) * time.Second)
		return
	}()
	return timer
}
