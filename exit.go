package main

import (
	"fmt"
	"os"
)

func main() {
	// defer will not be run when using os.Exit()
	defer fmt.Println("!")
	os.Exit(3)
}
