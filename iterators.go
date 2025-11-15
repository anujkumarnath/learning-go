package main

import (
	"fmt"
	"iter"
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
*/

func main() {
	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)

	for e := range lst.All() {
		fmt.Println(e)
	}
}
