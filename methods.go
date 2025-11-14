package main

import (
	"fmt"
)

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return  2 * (r.width + r.height)
}

func main() {
	r := rect{width: 10, height: 20}
	fmt.Println("area", r.area())
	fmt.Println("perimeter", r.perim())

	// go automatically handles conversion between
	// values and pointers for method calls
	rp := &r
	fmt.Println("area", rp.area())
	fmt.Println("perimeter", rp.perim())
}
