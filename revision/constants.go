package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	fmt.Println(s)

	const n = 500000000

	const d = 3e20 / n 
	fmt.Println(d)
	
	// numeric constants have no type unless given one explicitly
	fmt.Println(int64(d))

	// numeric constants can be given a type by using them in a context
	// that requires one, such as a variable assignment or function call
	// For example, math.Sin expects a float64
	fmt.Println(math.Sin(n))
}
