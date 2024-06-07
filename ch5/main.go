package main

import (
	"fmt"
)

func main() {

	result, err := MapElementCount("<html><body><p>hoehoge</p><div><div><p>piyo</p></div></div></body></html>")
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}
