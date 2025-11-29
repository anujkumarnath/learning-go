package main

import "fmt"
import "time"

func main() {
	timer1 := time.NewTimer(time.Second * 2)
	<-timer1.C

	fmt.Println("Timer 1 has fired")

	timer2 := time.NewTimer(time.Second)

	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	time.Sleep(2 * time.Second)
}
