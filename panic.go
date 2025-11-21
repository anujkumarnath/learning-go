package main

import (
	"os"
	"path/filepath"
)

// panics are to fail fast on errors that shouldn't occur
func main() {
	// comment it to have the file created
	panic("a problem")


	// panic are used when we don't know
	// how to handle and unexpected error
	path := filepath.Join(os.TempDir(), "file")
	_, err := os.Create(path)

	if err != nil {
		panic(err)
	}
}
