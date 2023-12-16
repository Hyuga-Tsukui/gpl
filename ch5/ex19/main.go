package main

import "fmt"

func main() {
	fmt.Println(noreturn())
}

func noreturn() (result int) {
	defer func() {
		if r := recover(); r != nil {
			result = r.(int)
		}
	}()
	panic(30)
}
