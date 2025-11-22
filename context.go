package main

import (
	"net/http"
	"fmt"
	"time"
)

func hello(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	fmt.Println("server: hello handler started")
	defer fmt.Println("server: hello handler ended")

	select {
	case <-time.After(10 * time.Second):
		fmt.Fprintf(w, "hello\n")
	case <-ctx.Done():
		err := ctx.Err()
		fmt.Println("server:", err)
		internalErr := http.StatusInternalServerError
		http.Error(w, err.Error(), internalErr)
	}
}

func main() {
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":3000", nil)
}
