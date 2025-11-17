package main

import "fmt"

func main() {
	done := make(chan bool)

	gen := func() <-chan int {
		out := make(chan int)

		go func() {
			num := 1
			for {
				select {
				case <-done:
					fmt.Println("--- GENERATION STOPPED ---")
					close(out)
					return
				default:
				}

				// Guranteed safe send (will never send after done)
				out <- num
				fmt.Printf("%-10s : %3d\n", "sent", num)
				num++
			}
		}()

		return out
	}

	for num := range gen() {
		fmt.Printf("%-10s : %3d\n", "received", num)	
		if num == 100 {
			close(done)
		}
	}
}
