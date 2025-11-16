package main

import (
	"fmt"
)

func SliceIndex[S ~[]E, E comparable](s S, e E) int {
	for i := range s {
		if e == s[i] {
			return i
		}
	}

	return -1
}

type List[T any] struct {
	head, tail *element[T]	
}

type element[T any] struct {
	val T
	next *element[T]
}

func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.tail = &element[T]{val: v}
		lst.head = lst.tail
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) AllElements() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)  	
	}
	return elems
}

func main() {
	var s = []string{"foo", "bar", "zoo"}	
	fmt.Println("index of zoo:", SliceIndex(s, "zoo"))
	fmt.Println("index of foo:", SliceIndex[[]string , string](s, "foo"))

	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)

	fmt.Println("list:", lst.AllElements())
}
