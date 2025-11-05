package main

import (
	"context"
	"fmt"
)

const Limit = 100

func main() {

	gen := func (ctx context.Context) <-chan int {
		dst := make(chan int)

		go func() {
			num := 1
			for {
				select {
				case <- ctx.Done():
					return
				case dst <- num:
					num++
					if num > Limit {
						/* The loop may sometimes be able to run an extra iteration
						 * before it receives signal from ctx.Done() channel, causing this leak
						 * It would not print from the consumer because the consumer loop breaks
						 * and doesn't read anymore
						 */
						fmt.Println("leaking", num)
					}
				}
			}
		}()

		return dst
	}

	ctx, cancel := context.WithCancel(context.Background())
	/*
	 * cancel ill be executed before exiting main and
	 * it will close the Done channel returned by the context.WithCancel call
	 */
	defer cancel()

	for num := range gen(ctx) {
		if num == Limit {
			fmt.Println(num)
			break
		}
	}

}
