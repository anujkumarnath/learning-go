package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup
	var ops atomic.Uint64
	var num int = 0

	for range 50 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range 1000 {
				// atomic add
				ops.Add(1)
				// normal add
				num++
			}
			fmt.Println("Worker done")
		}()
	}

	wg.Wait()

	fmt.Println("ops:", ops.Load(), "num:", num)
}
