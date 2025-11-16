package main

import (
	"fmt"
)

// fo is Go's only looping construct
func main() {
	i := 0
	for i <= 3 {
		fmt.Println(i)
		i++
	}
	
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	for i := range 10 {
		fmt.Println("range", i)
	}

	j := 1
	for {
		fmt.Println("loop")
		if j == 5 {
			break
		}
		j++
	}

	for i := range 6 {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}
}

