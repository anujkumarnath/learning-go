package main

import (
	"fmt"
	"path/filepath"
	"os"
)

// defer is similar to finally in other langauges
// it is meant for running some function later in a program's execution 
// usually for a purpose of cleanup
func main() {
	path := filepath.Join(os.TempDir(), "defer.txt")
	f := createFile(path)
	defer closeFile(f)
	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("creating")

	f, err := os.Create(p)

	if err != nil {
		panic(err)
	}

	return f
}

func writeFile(file *os.File) {
	fmt.Println("writing")
	fmt.Fprintf(file, "data")
}

func closeFile(file *os.File) {
	fmt.Println("closing")
	if err := file.Close(); err != nil {
		panic(err)
	}
}
