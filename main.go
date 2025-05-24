package main

import (
	"fmt"
	"os"
)

func main() {

	args := os.Args[1:]

	switch args[0] {

	case "add":
		Add(args[1:])
	case "update":
		Update(args[1:])
	case "delete":
		Delete(args[1:])
	case "mark-in-progress":
		MarkInProgress(args[1:])
	case "mark-done":
		MarkDone(args[1:])
	case "list":
		List(args[1:])
	case "help":
		Help()
	default:
		fmt.Printf("Invalid option: %v\n", args[0])
		Help()
	}

}
