package main

import (
	"fmt"
)

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num = %v", b.num)
}

type container struct {
	base
	str string
}

func main() {
	co := container {
		base: base{
			num: 2,
		},
		str: "some name",
	}

	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)
	fmt.Println("also num:", co.base.num)

	// since container embeds base,
	// the methods of base also becomes method of container
	fmt.Println("describe:", co.describe())

	type describer interface {
		describe() string
	}

	// by embedding base, container also implents the describer interface
	// which was implemented by base only
	// see below - no type issue with assignment (co is assignable to type describer) 
	var d describer = co
	fmt.Println("describer:", d.describe())
}
