package commands

import (
	"ModEd/hr/controller"
	"ModEd/hr/util"
	"flag"
	"fmt"

	"gorm.io/gorm"
)

// usage : go run hr/cli/HumanResourceCLI.go update student status -field="value"
// required field : id, status !!
func updateStudentStatus(args []string, tx *gorm.DB) error {
	fs := flag.NewFlagSet("updateStatus", flag.ExitOnError)
	studentID := fs.String("id", "", "Student ID to update status")
	status := fs.String("status", "", "New Status (ACTIVE, GRADUATED, or DROP)")
	fs.Parse(args)

	err := util.NewValidationChain(fs).
		Required("id").
		Length("id", 11).
		Required("status").
		Validate()
	if err != nil {
		fs.Usage()
		return fmt.Errorf("validation error: %v", err)
	}

	newStatus, err := util.StatusFromString(*status)
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}

	hrFacade := controller.NewHRFacade(tx)

	if err := hrFacade.UpdateStudentStatus(*studentID, newStatus); err != nil {
		return fmt.Errorf("failed to update student status: %v", err)
	}

	fmt.Printf("Student %s status successfully updated to %s!\n", *studentID, *status)
	return nil
}
