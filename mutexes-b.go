package main

import (
	"fmt"
	"sync"
)

// run go run -race mutexes-b.go to detect race condition
func main() {
	var i int
	var mu sync.Mutex
	var wg sync.WaitGroup

	incr := func(n int) {
		defer wg.Done()
		mu.Lock()
		defer mu.Unlock()
		for range n {
			i++	
		}
	}

	wg.Add(2)
	go incr(1000)
	go incr(1000)
	wg.Wait()

	fmt.Println("i:", i)
}
