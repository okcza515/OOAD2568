package commands

import (
	commonModel "ModEd/common/model"
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

	allowedStatuses := []string{
		util.StatusToString(commonModel.ACTIVE),
		util.StatusToString(commonModel.GRADUATED),
		util.StatusToString(commonModel.DROP),
	}

	validator := util.NewValidationChain(fs)
	validator.Field("id").Required().Length(11).Regex(`^[0-9]{11}$`)
	validator.Field("status").Required().AllowedValues(allowedStatuses)
	err := validator.Validate()
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
