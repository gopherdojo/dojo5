package timer

import (
	"time"
)

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
