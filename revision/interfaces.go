package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	height, width  float64
}

func (r *rect) area() float64 {
	return r.height * r.width
}

func (r *rect) perim() float64 {
	return 2 * r.height + 2 * r.width
}

type circle struct {
	radius float64
}

func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c *circle) perim() float64 {
	return 2 * math.Pi * c.radius 
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func detectCircle(g geometry) {
	if c, ok := g.(*circle); ok {
		fmt.Println("cirle with radius", c.radius)
	}

	switch t := g.(type) {
	case *circle:
		fmt.Println("detected circle:", t)
	case *rect:
		fmt.Println("detected rect:", t)
	default:
		fmt.Println("detected unknown")
	}
}

func main() {
	r := rect{10, 20}
	c := circle{4}

	fmt.Println(r.area())

	measure(&r)
	measure(&c)
	
	detectCircle(&r)
	detectCircle(&c)
}
