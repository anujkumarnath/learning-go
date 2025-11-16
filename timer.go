package main

import (
	"fmt"
	"time"
)

func main() {
	// this provides a blocking channel to wait on
	// that we are doing below
	// it will notify us through the channel after 2 seconds
	timer1 := time.NewTimer(2 * time.Second)

	// waiting for message from returned channel C
	<- timer1.C
	fmt.Println("Timer 1 fired")
	

	// the difference between using time.Sleep and
	// this technique is that the timer can be cancelled 
	// before it fires
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	// giving the timer2 enough time to fire, if it ever was going to,
	// to show it is infact stopped
	time.Sleep(2 * time.Second)
}
