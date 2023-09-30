package main

import (
	"fmt"
	"log"
	"os"

	github "github.com/arthurh0812/the-go-programming-language/4-composite-types/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatalf("Searching Issues failed: %s", err)
	}

	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d\t%9.9s\t%.55s\n", item.Number, item.User.Login, item.Title)
	}
}
