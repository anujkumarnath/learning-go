package main

import (
	"fmt"
	"net/http"
)

func main() {
	// http.HandleFunc binds the handler function we defined
	// and sets up the "default router" in the net/http package 
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	// start the server
	// putting nil in place of a router tells it to
	// use the default router we have just set above
	http.ListenAndServe(":3000", nil)
}

// a handler is an object implementing http.Header interface
// a function serving as a handler takes two arguments
// 1. http.ResponseWriter and
// 2. http.Request
func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}
