package main

import (
	"fmt"
)

func seq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	next := seq()
	fmt.Println(next())
	fmt.Println(next())

	next = seq()
	fmt.Println(next())
	fmt.Println(next())
}
