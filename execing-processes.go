package main

import (
	"os"
	"os/exec"
	"syscall"
)

// exec.Command is used when we need to access
// an external process from a running Go process
// exec replaces the current go process with another (perhans no-Go) process

func main() {
	// exec requires the absolute path to the binary
	// so we use exec.LookPath to find it (probably /bin/ls)
	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}
	
	// exec requires the arguments to be in the form of a slice
	// NOTE: the first arugment should be the program name
	args := []string{"ls", "-a", "-l", "-h"}

	// exec also requires some environment variables
	// we are now just passing the existing environment variables
	env := os.Environ()
	
	execErr := syscall.Exec(binary, args, env)
	
	// if there is an error we will get a return value
	// otherwise our prgram will be replaced by ls
	if execErr != nil {
		panic(execErr)
	}
}
