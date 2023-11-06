package main

import (
	"gpl/ch4/ex10"
	"log"
	"os"
)

func main() {
	_, result2, _, err := ex10.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	if err := ex10.Report.Execute(os.Stdout, result2); err != nil {
		log.Fatal(err)
	}
}
