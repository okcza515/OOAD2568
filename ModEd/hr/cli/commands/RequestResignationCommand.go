package commands

import (
	"ModEd/hr/controller"
	hrModel "ModEd/hr/model"
	"ModEd/hr/util"
	"flag"
	"fmt"

	"gorm.io/gorm"
)

// usage: go run hr/cli/HumanResourceCLI.go requestResignation -id="66050001" -reason="ย้ายคณะ"
func (c *RequestResignationCommand) Execute(args []string, tx *gorm.DB) error {
	fs := flag.NewFlagSet("requestResignation", flag.ExitOnError)
	studentID := fs.String("id", "", "Student ID")
	instructorID := fs.String("_id", "", "Instructor ID")
	reason := fs.String("reason", "", "Reason for resignation")
	role := fs.String("role", "", "Role of the requester (e.g., Student, Instructor)")
	fs.Parse(args)

	if err := util.ValidateRequiredFlags(fs, []string{"id", "reason", "role"}); err != nil {
		if err = util.ValidateRequiredFlags(fs, []string{"_id", "reason", "role"}); err != nil {
			fs.Usage()
			return fmt.Errorf("Validation error: %v\n", err)
		}
	}

	var requesterID string
	switch *role {
	case "Student":
		if *studentID == "" {
			return fmt.Errorf("Student role requires -studentID")
		}
		requesterID = *studentID
	case "Instructor":
		if *instructorID == "" {
			return fmt.Errorf("Instructor role requires -id")
		}
		requesterID = *instructorID
	default:
		return fmt.Errorf("Invalid role. Must be 'Student' or 'Instructor'")
	}

	db := util.OpenDatabase(*util.DatabasePath)
	hrFacade := controller.NewHRFacade(db)

	request := hrModel.NewRequestResignationBuilder().
		WithStudentID(requesterID).
		WithReason(*reason).
		Build()

	if err := hrFacade.SubmitResignationRequest(request); err != nil {
		return fmt.Errorf("Failed to submit resignation request: %v\n", err)
	}

	fmt.Println("Resignation request submitted successfully.")
	return nil
}
