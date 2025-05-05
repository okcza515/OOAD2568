package commands

import (
	"ModEd/hr/controller"
	"ModEd/hr/util"
	"flag"
	"fmt"

	"gorm.io/gorm"
)

type AddInstructorCommand struct{}

func (cmd *AddInstructorCommand) Execute(args []string, tx *gorm.DB) error {
	fs := flag.NewFlagSet("add-instructor", flag.ExitOnError)
	instructorCode := fs.String("code", "", "Instructor Code")
	firstName := fs.String("fname", "", "First Name")¯
	lastName := fs.String("lname", "", "Last Name")
	email := fs.String("email", "", "Email")
	startDate := fs.String("startDate", "", "Start Date")
	department := fs.String("department", "", "Department")
	gender := fs.String("gender", "", "Gender")
	citizenID := fs.String("citizenID", "", "Citizen ID")
	phoneNumber := fs.String("phone", "", "Phone Number")
	salary := fs.Int("salary", 0, "Salary")
	academicPos := fs.String("academicPos", "", "Academic Position")
	departmentPos := fs.String("departmentPos", "", "Department Position")

	fs.Parse(args)
	validator := util.NewValidationChain(fs)
	validator.Field("code").Required().IsInstructorID()
	validator.Field("fname").Required()
	validator.Field("lname").Required()
	validator.Field("email").Required().IsEmail()
	validator.Field("startDate").Required().IsDate()
	validator.Field("department").Required()
	validator.Field("gender").Required()
	validator.Field("citizenID").Required().Length(13)
	validator.Field("phone").Required()
	validator.Field("salary").Required()
	validator.Field("academicPos").Required()
	validator.Field("departmentPos").Required()
	err := validator.Validate()
	if err != nil {
		fs.Usage()
		return fmt.Errorf("validation error: %v", err)
	}

	// TODO: Implement Add Instructor logic
	err = controller.AddInstructor(
		tx,
		*instructorCode,
		*firstName,
		*lastName,
		*email,
		*startDate,
		*department,
		*gender,
		*citizenID,
		*phoneNumber,
		*salary,
		*academicPos,
		*departmentPos,
	)

	if err != nil {
		return fmt.Errorf("transaction failed: %v", err)
	}
	fmt.Println("Instructor added and HR info updated successfully!")
	return nil
}
