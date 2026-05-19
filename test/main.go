package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("http://example.com")
	if err != nil {
		fmt.Println("An error occured", err)
		return
	}

	fmt.Println(resp)
}