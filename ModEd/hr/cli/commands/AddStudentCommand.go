package commands

import (
	"ModEd/hr/controller"
	"ModEd/hr/util"
	"flag"
	"fmt"

	"gorm.io/gorm"
)

type AddStudentCommand struct{}

func (cmd *AddStudentCommand) Execute(args []string, tx *gorm.DB) error {
	fs := flag.NewFlagSet("add-student", flag.ExitOnError)
	studentID := fs.String("id", "", "Student ID")
	firstName := fs.String("fname", "", "First Name")
	lastName := fs.String("lname", "", "Last Name")
	email := fs.String("email", "", "Email")
	startDate := fs.String("startDate", "", "Start Date")
	birthDate := fs.String("birthDate", "", "Birth Date")
	program := fs.String("program", "", "Program")
	department := fs.String("department", "", "Department")
	status := fs.String("status", "", "Status")
	gender := fs.String("gender", "", "Gender")
	citizenID := fs.String("citizenID", "", "Citizen ID")
	phoneNumber := fs.String("phone", "", "Phone Number")
	fs.Parse(args)

	validator := util.NewValidationChain(fs)
	validator.Field("id").Required().IsStudentID()
	validator.Field("fname").Required()
	validator.Field("lname").Required()
	validator.Field("email").Required().IsEmail()
	validator.Field("gender").Required().AllowedValues([]string{"Male", "Female", "Other"})
	validator.Field("citizenID").Required().Length(13)
	validator.Field("phone").Required()
	err := validator.Validate()
	if err != nil {
		fs.Usage()
		return fmt.Errorf("validation error: %v", err)
	}

	studentController := controller.NewStudentHRController(tx)
	err = studentController.AddStudent(
		*studentID,
		*firstName,
		*lastName,
		*email,
		*startDate,
		*birthDate,
		*program,
		*department,
		*status,
		*gender,
		*citizenID,
		*phoneNumber,
	)

	if err != nil {
		return fmt.Errorf("transaction failed: %v", err)
	}

	fmt.Println("Student added and HR info updated successfully!")

	return nil
}
