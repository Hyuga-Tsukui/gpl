package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("https://pkg.go.dev/net/http")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	result, err := MapElementCount(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}
