package main

import (
	"fmt"
	"math"
)

// interfaces are named collection of method signatures
// named collection: 'gerometry'
type geometry interface {
	// function signature
	area()  float64
	// function signature
	perim() float64
}

// implementation on rect type
type rect struct {
	width, height float64
}

// implementation on circle type
type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.height * r.width
}

func (r rect) perim() float64 {
	return 2 * (r.height + r.width)
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func detectCircle(g geometry) {
	// type assestion
	if c, ok := g.(circle); ok {
		fmt.Println("circle with radius", c.radius)
	}

	// vs type switch
	switch g.(type) {
		case circle:
			fmt.Println("the type of the interface is circle")
		case rect:
			fmt.Println("the type of the interface is rect")
	}
}

func main() {
	r := rect{10, 20}	
	c := circle{4}

	measure(r)
	measure(c)

	detectCircle(r)
	detectCircle(c)
}
