package main

import "fmt"

type person struct {
	name string
	age int
}

func newPerson(name string) *person {
	p := person{ name: name, age: 42 }
	return &p
}

func main() {
	fmt.Println(person{"Bob", 20})

	fmt.Println(person{name: "Alice", age: 30})

	// omitted fields are zero valued
	fmt.Println(person{name: "Fred"})

	fmt.Println(&person{"Ann", 34})

	// idiomatic to encapsulate new struct creation in concstuctor functions
	fmt.Println(newPerson("John"))

	s := person{name: "Sean", age: 50}
	// dot to access struct members
	fmt.Println(s.name)

	sp := &s
	// can also use dot with struct pointer,
	// struct pointers are dereferenced automatically
	fmt.Println(sp.age)

	// structs are mutable
	sp.age = 51
	fmt.Println(sp.age)

	// anonymours struct type
	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}

	fmt.Println(dog)

}
