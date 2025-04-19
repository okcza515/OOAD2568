package commands

import (
	"ModEd/hr/controller"
	"ModEd/hr/util"
	"flag"
	"fmt"
	"os"
)

// usage : go run hr/cli/HumanResourceCLI.go updateStatus -field="value"
// required field : id, status !!
func (c *UpdateStudentStatusCommand) Run(args []string){
	fs := flag.NewFlagSet("updateStatus", flag.ExitOnError)
	studentID := fs.String("id", "", "Student ID to update status")
	status := fs.String("status", "", "New Status (ACTIVE, GRADUATED, or DROP)")
	fs.Parse(args)

	if err := util.ValidateRequiredFlags(fs, []string{"id", "status"}); err != nil {
        fmt.Printf("Validation error: %v\n", err)
        fs.Usage()
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
