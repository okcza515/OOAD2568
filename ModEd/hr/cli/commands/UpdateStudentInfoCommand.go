package commands

import (
	"ModEd/hr/controller"
	"ModEd/hr/util"
	"flag"
	"fmt"

	"gorm.io/gorm"
)

// usage : go run hr/cli/HumanResourceCLI.go update student info -field="value"
// required field : id !!

func updateStudentInfo(args []string, tx *gorm.DB) error {
	fs := flag.NewFlagSet("info", flag.ExitOnError)
	studentID := fs.String("id", "", "Student ID to update")
	firstName := fs.String("fname", "", "New first name")
	lastName := fs.String("lname", "", "New last name")
	gender := fs.String("gender", "", "New gender")
	citizenID := fs.String("citizenID", "", "New citizen ID")
	phoneNumber := fs.String("phone", "", "New phone number")
	email := fs.String("email", "", "New email")
	fs.Parse(args)

	validator := util.NewValidationChain(fs)
	validator.Field("id").Required().Length(11).Regex(`^[0-9]{11}$`)
	err := validator.Validate()
	if err != nil {
		fs.Usage()
		return fmt.Errorf("validation error: %v", err)
	}

	if err := controller.UpdateStudentInfo(tx, *studentID, *firstName, *lastName, *gender, *citizenID, *phoneNumber, *email); err != nil {
		return fmt.Errorf("failed to update student info: %v", err)
	}

	fmt.Println("Student updated successfully!")
	return nil
}
