package main

import (
	"ModEd/hr/controller"
	"fmt"
	"os"
)

func handleStudentCommand(hrController *controller.HRController, command string, args []string) {
	switch command {
	case "list":
		// listStudents(hrController, args)
	case "update":
		// updateStudent(hrController, args)
	case "add":
		// addStudent(hrController, args)
	case "delete":
		// deleteStudent(hrController, args)
	case "updateStatus":
		// updateStudentStatus(hrController, args)
	case "import":
		// importStudents(hrController, args)
	case "sync":
		// synchronizeStudents(hrController)
	default:
		fmt.Printf("Unknown student command: %s\n", command)
		fmt.Println("Available student commands: list, add, update, delete, updateStatus, import, sync")
		os.Exit(1)
	}
}
