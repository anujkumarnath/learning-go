package main

import (
	"fmt"
	"sync"
)

func main() {
	var i int // normal variable
	var mu sync.Mutex
	var wg sync.WaitGroup

	for range 10 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()
			for range 100 {
				i++
			}
		}()
	}

	wg.Wait()
	fmt.Println("i", i)
}
