package main

import (
	"fmt"
)

func fact(num int) int {
	if num <= 1 {
		return 1
	}
	return num * fact(num - 1)
}

func main() {
	num := 10
	fmt.Printf("fact(%d) = %d\n", num, fact(num))


	var fib func(n int) int
	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n - 1) + fib(n - 2)
	}
	
	n := 0
	fmt.Printf("fib(%d) = %d\n", n, fib(n))
}
