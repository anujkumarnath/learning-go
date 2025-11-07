package main

import (
	"fmt"
)

func main() {
	fmt.Println("3 + 4 =", add(3, 4))
	fmt.Println("1 + 2 + 3 + 4 + 0.3 =", fourAdd(1, 2, 3, 4, "Hi"))
}

func add(a int, b int) int {
	return a + b
}

// if multiple consecutive arguments have same type,
// we can only specify the type after the last one
func fourAdd(a, b, c, d int, str string) int {
	fmt.Println(str)
	return a + b + c + d
}
