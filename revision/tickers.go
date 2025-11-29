package main

import "fmt"
import "time"

func main() {
	ticker := time.NewTicker(time.Second)
	done := make(chan bool)

	go func() {
		for {
			select {
			case t := <-ticker.C:
				fmt.Println(t)
			case <-done:
				ticker.Stop()
				return
			}
		}
	}()

	time.Sleep(6 * time.Second)
	close(done)
	fmt.Println("Ticker stopped")
}
