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

	if err := util.ValidateRequiredFlags(fs, []string{"id", "status"}); err != nil {
		fs.Usage()
		return fmt.Errorf("Validation error: %v\n", err)
	}

	newStatus, err := util.StatusFromString(*status)
	if err != nil {
		return fmt.Errorf("Error: %v\n", err)
	}

	db := util.OpenDatabase(*util.DatabasePath)
	hrFacade := controller.NewHRFacade(db)

	if err := hrFacade.UpdateStudentStatus(*studentID, newStatus); err != nil {
		return fmt.Errorf("Failed to update student status: %v\n", err)
	}

	fmt.Printf("Student %s status successfully updated to %s!\n", *studentID, *status)
	return nil
}
