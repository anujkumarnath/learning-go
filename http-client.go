package main

import (
	"fmt"
	"bufio"
	"net/http"
)

func main() {
	res, err := http.Get("https://gobyexample.com")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	fmt.Println("Response status:", res.Status)

	scanner := bufio.NewScanner(res.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
