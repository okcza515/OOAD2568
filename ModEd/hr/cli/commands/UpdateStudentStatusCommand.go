package commands

import (
	"ModEd/hr/controller"
	"ModEd/hr/util"
	"flag"
	"fmt"
	"os"
)

func (c *UpdateStudentStatusCommand) Run(args []string){
	fs := flag.NewFlagSet("updateStatus", flag.ExitOnError)
	studentID := fs.String("id", "", "Student ID to update status")
	status := fs.String("status", "", "New Status (ACTIVE, GRADUATED, or DROP)")
	fs.Parse(args)

	if *studentID == "" || *status == "" {
		fmt.Println("Error: Student ID and Status are required.")
		fmt.Println("Usage: go run humanresourcecli.go [-database=<path>] updateStatus -id=<studentID> -status=<ACTIVE|GRADUATED|DROP>")
		os.Exit(1)
	}


	newStatus, err := util.StatusFromString(*status)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	db := util.OpenDatabase(*util.DatabasePath)
	hrFacade := controller.NewHRFacade(db)

	if err := hrFacade.UpdateStudentStatus(*studentID, newStatus); err != nil {
		fmt.Printf("Failed to update student status: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Student %s status successfully updated to %s!\n", *studentID, *status)
}
