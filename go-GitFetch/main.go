package main

import (
	"flag"
	"fmt"
	"go-GitFetch/git"
	"go-GitFetch/utils"
	"log"
	"os"
)

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) != 1 {
		fmt.Println("Please provide exactly one argument: the GitHub username.")
		os.Exit(1)
	}

	username := args[0]

	events, err := git.GetEvents(username)

	if err != nil {
		log.Fatalf("Error fetching events %v", err)
	}

	if len(events) == 0 {
		fmt.Printf("No events found for user: %s\n", username)
		return
	}

	for _, event := range events {
		fmt.Println(utils.FormatEvent(event))
	}
}
