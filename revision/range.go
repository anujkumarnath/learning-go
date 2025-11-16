package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3}
	sum := 0
	for _, n := range nums {
		sum += n
	}
	fmt.Println("sum:", sum)

	for i, n := range nums {
		if n == 3 {
			fmt.Println("index:", i)
		}
	}
	
	kvs := map[string]int{"one": 1, "two": 2, "three": 3}
	for k, v := range kvs {
		fmt.Printf("%5s -> %d\n", k, v)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
