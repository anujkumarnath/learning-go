package main

import (
	"fmt"
)

func ival(i int) {
	i = 0
}

func iptr(i *int) {
	*i = 0
}

func main() {
	i := 1
	fmt.Println("i:", i)

	ival(i)
	fmt.Println("i after ival(i):", i)

	iptr(&i)
	fmt.Println("i after iptr(i):", i)

	fmt.Println("ptr:", &i)
}
