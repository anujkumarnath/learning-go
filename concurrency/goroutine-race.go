package main

import "fmt"

func main() {
	done := make(chan struct{})

	gen := func() <-chan int {
		dst := make(chan int)
		go func() {
			num := 1
			for {
				select {
				case <- done:
					close(dst)
					return
				case dst <- num:
					num++
				}
			}

		}()
		return dst
	}

	ch := gen();

	for num := range ch {
		fmt.Println(num)
		if num == 100 {
			break;
		}
	}

	close(done)

	/* By the time the goroutine receives the done signal,
	 * another iteration of for-select may run, and dst <- num may get
	 * So we may sometimes see prints till 101 and sometimes till 100
	 * executed, even if the consumer is consuming - we won't see the print
	 * in above loop. But the producer goroutine will remain blocked because
	 * there is no one to consume - the goroutine will leak (okay for main funciton - program will exit)
	 * But to free the goroutine completely - we need to drain the value as below
	 */
	for num := range ch {
		// Just printing to verify, the print is not required for draining
		fmt.Println(num)
	}
}
