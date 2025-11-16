package main

import (
	"fmt"
)

type rect struct {
	width  int
	height int
}

func (r *rect) Area() int {
	return r.width * r.height
}

func main() {
	r := rect{height: 10, width: 20}
	fmt.Println(r.Area())

	rp := &r
	fmt.Println(rp.Area())
}
