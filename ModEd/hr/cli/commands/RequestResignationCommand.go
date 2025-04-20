package commands

import (
	"ModEd/hr/controller"
	hrModel "ModEd/hr/model"
	"ModEd/hr/util"
	"flag"
	"fmt"
	"os"
)


// usage: go run hr/cli/HumanResourceCLI.go requestResignation -id="66050001" -reason="ย้ายคณะ"
func (c *RequestResignationCommand) Run(args []string) {
	fs := flag.NewFlagSet("requestResignation", flag.ExitOnError)
	studentID := fs.String("id", "", "Student ID")
	instructorID := fs.String("_id", "", "Instructor ID")
	reason := fs.String("reason", "", "Reason for resignation")
	role := fs.String("role", "", "Role of the requester (e.g., Student, Instructor)")
	fs.Parse(args)

	if err := (util.ValidateRequiredFlags(fs, []string{"id","_id","reason", "role"})||util.ValidateRequiredFlags(fs, []string{"id","_id","reason", "role"})); err != nil {
		fmt.Printf("Validation error: %v\n", err)
		fs.Usage()
		os.Exit(1)
	}

	var requesterID string
	switch *role {
	case "Student":
		if *studentID == "" {
			fmt.Println("Student role requires -studentID")
			os.Exit(1)
		}
		requesterID = *studentID
	case "Instructor":
		if *instructorID == "" {
			fmt.Println("Instructor role requires -id")
			os.Exit(1)
		}
		requesterID = *instructorID
	default:
		fmt.Println("Invalid role. Must be 'Student' or 'Instructor'")
		os.Exit(1)
	}

	db := util.OpenDatabase(*util.DatabasePath)
	hrFacade := controller.NewHRFacade(db)

	request := hrModel.NewRequestResignationBuilder().
		WithStudentID(requesterID). 
		WithReason(*reason).
		Build()

	if err := hrFacade.SubmitResignationRequest(request); err != nil {
		fmt.Printf("Failed to submit resignation request: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Resignation request submitted successfully.")
}
