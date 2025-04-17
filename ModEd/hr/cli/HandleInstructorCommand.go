package main

import (
	"fmt"
	"os"

	"ModEd/hr/controller"
)

func handleInstructorCommand(hrController *controller.HRController, command string, args []string) {
	switch command {
	case "list":
		// listInstructors(hrController, args)
	case "add":
		// addInstructor(hrController, args)
	case "update":
		// updateInstructor(hrController, args)
	case "delete":
		// deleteInstructor(hrController, args)
	default:
		fmt.Printf("Unknown instructor command: %s\n", command)
		fmt.Println("Available instructor commands: list, add, update, delete, etc.")
		os.Exit(1)
	}
}
