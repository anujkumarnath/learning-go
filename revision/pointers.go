package main

import (
	"fmt"
)

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	num := 1
	zeroval(num)
	fmt.Println(num)

	var ptr *int
	ptr = &num

	zeroptr(ptr)
	fmt.Println(num)
}
