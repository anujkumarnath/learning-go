package greetings

import "fmt"

// a function whose name starts with a capital letter
// can be called from code not belonging to the same package
func Hello(name string) string {
	message := fmt.Sprintf("Hi %v! Welcomei!", name)
	return message
}
