package main

import (
	"fmt"
	"os/exec"
	"errors"
	"io"
)

func main() {
	dateCmd := exec.Command("date")	

	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("date >")
	fmt.Println(string(dateOut))

	_, err = exec.Command("date", "-x").Output()
	if err != nil {
		// if there is any error executing the program
		var execErr *exec.Error
		// if the program ran but exited with a non-zero code
		var exitErr *exec.ExitError

		switch {
		case errors.As(err, &execErr):
			fmt.Println("failed executing:", err)
		case errors.As(err, &exitErr):
			exitCode := exitErr.ExitCode()
			fmt.Println("command exit rc =", exitCode)
		default:
			panic(err)
		}
	}

	grepCmd := exec.Command("grep", "hello")
	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()

	grepCmd.Start()
	grepIn.Write([]byte("hellobro grep\n goodbye grep"))
	grepIn.Close()
	
	grepBytes, _ := io.ReadAll(grepOut)
	// necessary for finalizing the process
	// without it the process will become a zombie process,
	// and such processes may pile up in system
	// if we don't use it we will miss the error
	// If stderr has data and we don't read it, the child can block forever unless we Wait()
	// pipes may stay open unless GC closes them
	grepCmd.Wait()

	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))

	// till now we need to explicitly pass the command and argument array
	// to pass the entire command in a single string, we can use the bash -c option
	lsCmd := exec.Command("bash", "-c", "ls -l -a -h")
	lsOut, err := lsCmd.Output()
	if err !=  nil {
		panic(err)
	}
	fmt.Println("> ls -l -a -h")
	fmt.Println(string(lsOut))
}
