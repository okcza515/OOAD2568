package commands

import (
	"ModEd/hr/controller"
	"ModEd/hr/util"
	"flag"
	"fmt"

	"gorm.io/gorm"
)

type UpdateStudentInfoCommand struct{}

// usage : go run hr/cli/HumanResourceCLI.go update student info -field="value"
// required field : id !!

func (cmd *UpdateStudentInfoCommand) Execute(args []string, tx *gorm.DB) error {
	fs := flag.NewFlagSet("update-student-info", flag.ExitOnError)
	studentID := fs.String("id", "", "Student ID")
	firstName := fs.String("fname", "", "First Name")
	lastName := fs.String("lname", "", "Last Name")
	email := fs.String("email", "", "Email")
	gender := fs.String("gender", "", "Gender")
	citizenID := fs.String("citizenID", "", "Citizen ID")
	phoneNumber := fs.String("phone", "", "Phone Number")
	fs.Parse(args)

	validator := util.NewValidationChain(fs)
	validator.Field("id").Required().Length(11).Regex(`^[0-9]{11}$`)
	err := validator.Validate()
	if err != nil {
		fs.Usage()
		return fmt.Errorf("validation error: %v", err)
	}
	studentController := controller.NewStudentHRController(tx)
	if err := studentController.UpdateStudentInfo(
		*studentID,
		*firstName,
		*lastName,
		*email,
		*gender,
		*citizenID,
		*phoneNumber,
	); err != nil {
		return fmt.Errorf("failed to update student info: %v", err)
	}

	fmt.Println("Student updated successfully!")
	return nil
}
