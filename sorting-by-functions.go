package main

import (
	"fmt"
	"slices"
	"cmp"
)

func main() {
	fruits := []string{"peach", "banana", "kiwi"}
	lenCmp := func(a, b string) int {
		return cmp.Compare(len(a), len(b))
	}

	slices.SortFunc(fruits, lenCmp)
	fmt.Println(fruits)

	type Person struct {
		name string
		age  int
	}

	people := []Person{
		Person{"Jax", 37},
		Person{"TJ", 25},
		Person{name: "Alex", age: 72},
		Person{name: "Bob", age: 50},
		Person{name: "John", age: 13},
		Person{name: "Paul", age: 34},
		
		// this also works
		// {"Jax", 37},
		// {"TJ", 25},
		// {name: "Alex", age: 72},
		// {name: "Bob", age: 50},
		// {name: "John", age: 13},
		// {name: "Paul", age: 34},
	}

	slices.SortFunc(people,
		func(p1, p2 Person) int {
			return cmp.Compare(p1.age, p2.age)
		})

	fmt.Println("sorted by value", people)

	peoplePointer := []*Person{
		&Person{"Jax", 37},
		&Person{"TJ", 25},
		&Person{name: "Alex", age: 72},
		&Person{name: "Bob", age: 50},
		&Person{name: "John", age: 13},
		&Person{name: "Paul", age: 34},
	}

	slices.SortFunc(peoplePointer,
		func(p1, p2 *Person) int {
			return cmp.Compare(p1.age, p2.age)
		})

	fmt.Print("sorted by pointer:")
	for _, personPointer := range peoplePointer {
		fmt.Print(*personPointer, ", ")
	}
	fmt.Println()
}
