package commands

import (
	"flag"
	"fmt"

	"ModEd/hr/controller"
	"ModEd/hr/util"

	"gorm.io/gorm"
)

type UpdateInstructorInfoCommand struct{}

// usage: go run hr/cli/HumanResourceCLI.go update instructor info -id="66050001" -field="position" -value="Professor"
func (cmd *UpdateInstructorInfoCommand) Execute(args []string, tx *gorm.DB) error {
	fs := flag.NewFlagSet("update-instructor-info", flag.ExitOnError)
	instructorID := fs.String("id", "", "Instructor ID to update")
	firstName := fs.String("fname", "", "Instructor's first name")
	lastName := fs.String("lname", "", "Instructor's last name")
	email := fs.String("email", "", "Instructor's email")
	gender := fs.String("gender", "", "Instructor's gender")
	citizenID := fs.String("citizenID", "", "Instructor's citizen ID")
	phoneNumber := fs.String("phone", "", "Instructor's phone number")
	academicPos := fs.String("academicPos", "", "Instructor's academic position")
	departmentPos := fs.String("departmentPos", "", "Instructor's department position")
	fs.Parse(args)

	validator := util.NewValidationChain(fs)
	validator.Field("id").Required().Length(11).Regex(`^[0-9]{11}$`)
	validator.Field("field").Required()
	validator.Field("value").Required()
	err := validator.Validate()
	if err != nil {
		fs.Usage()
		return fmt.Errorf("validation error: %v", err)
	}

	instructorController := controller.NewInstructorHRController(tx)
	if err := instructorController.UpdateInstructorInfo(
		*instructorID,
		*firstName,
		*lastName,
		*email,
		*gender,
		*citizenID,
		*phoneNumber,
		*academicPos,
		*departmentPos,
	)
	err != nil {
		return fmt.Errorf("failed to update instructor info: %v", err)
	}

	fmt.Println("Instructor updated successfully!")
	return nil
}
