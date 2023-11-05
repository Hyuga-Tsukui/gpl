package main

import (
	"fmt"
	"gpl/ch4/ex10"
	"log"
	"os"
)

func main() {
	result1, result2, result3, err := ex10.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issue1:\n", result1.TotalCount)
	fmt.Printf("%d issue2:\n", result2.TotalCount)
	fmt.Printf("%d issue3:\n", result3.TotalCount)
}
