package main

import (
	"fmt"
)

// generics are also called type parrameters
// type parrameter are written inside '[' & ']'
// here S ~[]E, E and comparable are type parameters
// the function expects types S and E in it's parameters
// and in type parameters we define what S and E are:
// S is a slice of E's, and the comparable contstraint means
// E's can be compared with ==/!= operators
func SliceIndex[S ~[]E, E comparable](s S, v E) int {
	for i := range s {
		if s[i] == v { // using E any will not allow this
			return i
		}
	}

	return -1
}

// NOTE:
// Type contstraints: any, comparable
// 
// T any, T comparable
// any        : no contstraint
// comparable : the elements should be comparable
// - must support comparison with == and != operators

// definition of a linked-list of generic type
// T any means the linked-list supports any type of elements
type List[T any] struct {
	head, tail *element[T]
}


// individual element(node) of the linked-list
// T any means the node works for any data type
// it stores the value and the pointer of 
type element[T any] struct {
	next *element[T]
	val T
}

/*
// Equivalent linked-list struct for int type
type List struct {
	head, tail *element
}

// Equivalent element struct for int values
type element struct {
	next *element
	val int
}
*/

// defining methods on generic type
func (list *List[T]) Push(v T) {
	if list.tail == nil {
		list.head = &element[T]{val: v}
		list.tail = list.head
	} else {
		list.tail.next = &element[T]{val: v}
		list.tail = list.tail.next
	}
}

func (list *List[T]) AllElements() []T {
	var elements []T
	for e := list.head; e != nil; e = e.next {
		elements = append(elements, e.val)
	} 
	return  elements
}

func main() {
	var s = []string{	"foo", "bar", "zoo"}
	// when invoking generic functions, we can avoid specifying types
	// by relying on type inference - don't have to specify typs S and E
	// while calling SliceIndex
	fmt.Println("index of zoo:", SliceIndex(s, "zoo"))

	// but we can specify them explicitly
	_ = SliceIndex[[]string, string](s, "zoo")

	list := List[int]{}
	list.Push(10)
	list.Push(13)
	list.Push(23)
	fmt.Println("list:", list.AllElements())
}
