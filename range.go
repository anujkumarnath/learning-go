package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 4, 5}
	sum := 0
	
	for _, num := range nums {
		sum += num
	}

	fmt.Println("sum", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index", i)
		}
	}

	kvs := map[string]int{"one": 1, "two": 2, "three": 3}

	for k, v := range kvs {
		fmt.Printf("kvs[%s]: %d\n", k, v)
	}

	for k := range kvs {
		fmt.Println("key:", k)
	}

	// ranging over unicode code points in a string
	for i, char := range "A long string" {
		fmt.Println(i, char)
	}
}
