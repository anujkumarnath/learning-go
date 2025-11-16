package main

import (
	"fmt"
)

func add(a, b int) int {
	return a+b
}

func vals() (int, int) {
	return 3, 7
}

func variadic(a, b int, strings ...string) {
	fmt.Println("a:", a, "b:", b)
	fmt.Println("strings:", strings)
}

func main() {
	fmt.Println(add(2, 3))
	fmt.Println(vals())
	variadic(1, 2, "bob", "alice")
	variadic(3, 4, []string{"john", "doe"}...)
}

