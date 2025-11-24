package greetings

import (
	"fmt"
	"errors"
)

// a function whose name starts with a capital letter
// can be called from code not belonging to the same package
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf("Hi %v! Welcome!", name)
	return message, nil
}
