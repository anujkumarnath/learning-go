package main

import (
	"fmt"
	"iter"
	"slices"
)

type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	val T
	next *element[T]
}

func (list *List[T]) Push(v T) {
	if list.tail == nil {
		list.head = &element[T]{val: v}
		list.tail = list.head
	} else {
		list.tail.next = &element[T]{val: v}
		list.tail = list.tail.next
	}
}

// can iterate over a generic type with range by
// implementing an iterator of type iter.Seq
func (list *List[T]) All() iter.Seq[T] {
	return func (yield func(T) bool) {
		for e := list.head; e != nil; e = e.next {
			if !yield(e.val) {
				return
			}
		}
	}
}

/* EXPLANATION:
// NOTE: go iterators are just syntactic sugar, unlike JS, Java, Python
// iterators in go are functions with special signature
// thery are only based on compiler generated yield callbacks
// All() returns this function
var it = func (yield func(T) bool) {
	for e := list.head; e != nil; e = e.next {
		if !yield(e.val) {
			return
		}
	}
}

// compiler creates a inline yeild function
// and converts the range call to
var yield = func(val int) bool {
	fmt.Println(val)
	return true
}

// it then creates a call like this
// and the iterator runs
// calling the yield with e.val inside the loop
// since it returns true everytime, the entire linked-list is traversed
it(yield)
*/

func genFib() iter.Seq[int] {
	return func(yield func(int) bool) {
		a, b := 1, 1
		for {
			if !yield(a) {
				return
			}
			a, b = b, a+b
		}
	}
}

func main() {
	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)

	for e := range lst.All() {
		fmt.Println(e)
	}

	// slices.Collect method takes an iterator and
	// collects all its value into a slice and rerturns it
	all := slices.Collect(lst.All())
	fmt.Println(all)

	// once the range loop hits break or returns early,
	// the yield function returns false
	for n := range genFib() {
		if n >= 10 {
			break
		}
		fmt.Println(n)
	}

	/*
	// for earyly break or return,
	// the compiler creates a return statemnt in the yield function
	yield := func(n int) bool {
		if n >= 10 {
			return false // compiler uses yield's return value to implement break
		}
		fmt.Println(n)
		return true
	}
	*/
}
