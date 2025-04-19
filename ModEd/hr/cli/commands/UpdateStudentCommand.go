package commands

import (
	"ModEd/hr/controller"
	hrModel "ModEd/hr/model"
	hrUtil "ModEd/hr/util"
	"flag"
	"fmt"
	"os"
)

// usage : go run hr/cli/HumanResourceCLI.go update -field="value"
// required field : id !!

func (c *UpdateStudentCommand) Run(args []string) {
	fs := flag.NewFlagSet("update", flag.ExitOnError)
	studentID := fs.String("id", "", "Student ID to update")
	firstName := fs.String("fname", "", "New First Name value")
	lastName := fs.String("lname", "", "New Last Name value")
	gender := fs.String("gender", "", "New Gender value")
	citizenID := fs.String("citizenID", "", "New Citizen ID value")
	phoneNumber := fs.String("phone", "", "New Phone Number value")
	emailStudent := fs.String("email", "", "New Email value")
	fs.Parse(args)

	if err := hrUtil.ValidateRequiredFlags(fs, []string{"id"}); err != nil {
		fmt.Printf("Validation error: %v\n", err)
		fs.Usage()
		os.Exit(1)
	}

	db := hrUtil.OpenDatabase(*hrUtil.DatabasePath)

	hrFacade := controller.NewHRFacade(db)
	studentInfo, err := hrFacade.GetStudentById(*studentID)
	if err != nil {
		fmt.Printf("Error retrieving student with ID %s: %v\n", *studentID, err)
		os.Exit(1)
	}

	studentInfoBuilder := hrModel.NewStudentInfoBuilder().
		WithFirstName(ifNotEmpty(*firstName, studentInfo.FirstName)).
		WithLastName(ifNotEmpty(*lastName, studentInfo.LastName)).
		WithGender(ifNotEmpty(*gender, studentInfo.Gender)).
		WithCitizenID(ifNotEmpty(*citizenID, studentInfo.CitizenID)).
		WithPhoneNumber(ifNotEmpty(*phoneNumber, studentInfo.PhoneNumber)).
		WithEmail(ifNotEmpty(*emailStudent, studentInfo.Email)).
		Build()

	if err := hrFacade.UpdateStudent(studentInfoBuilder); err != nil {
		fmt.Printf("Failed to update student info: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Student updated successfully!")
}

func ifNotEmpty(newValue, fallbackValue string) string {
	if newValue != "" {
		return newValue
	}
	return fallbackValue
}
