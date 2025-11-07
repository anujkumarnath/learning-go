package main

import (
	"fmt"
)

func intSeq() func() int {
	i := 0

	// the anonymous function closes over the
	// the variable i used from the parent funcion's context
	// forming a colsure
	return func() int {
		i++
		return i
	}
}

func main() {
	nextInt := intSeq();

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// state is unique to every returned function (nextInt)
	nextInt = intSeq();
	fmt.Println(nextInt())
}
