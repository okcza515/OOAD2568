package main

import (
	"ModEd/hr/controller"
	"flag"
	"fmt"
	"os"
)

var (
	databasePath = flag.String("database", "data/ModEd.bin", "Path of SQLite Database")
)

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) < 1 {
		fmt.Println("Usage: go run humanresourcecli.go {student|instructor} {list|add|update|delete|updateStatus|import|sync} [options]")
		os.Exit(1)
	}

	entity := args[0]
	command := args[1]
	commandArgs := args[2:]

	hrController := controller.NewHRController(*databasePath)

	switch entity {
	case "student":
		handleStudentCommand(hrController, command, commandArgs)
	case "instructor":
		handleInstructorCommand(hrController, command, commandArgs)
	default:
		fmt.Printf("Unknown entity: %s\n", entity)
		fmt.Println("Available entities: student, instructor")
		os.Exit(1)
	}
}
