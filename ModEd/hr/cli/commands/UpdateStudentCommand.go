package commands

import (
	commonController "ModEd/common/controller"
	"ModEd/hr/controller"
	hrModel "ModEd/hr/model"
	"ModEd/hr/util"
	hrUtil "ModEd/hr/util"
	"flag"
	"fmt"

	"gorm.io/gorm"
)

// usage : go run hr/cli/HumanResourceCLI.go update -field="value"
// required field : id !!

func (c *UpdateStudentCommand) Execute(args []string, tx *gorm.DB) error {
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
		fs.Usage()
		return fmt.Errorf("Validation error: %v\n", err)
	}

	db := hrUtil.OpenDatabase(*hrUtil.DatabasePath)

	// Create a TransactionManager instance.
	tm := &util.TransactionManager{DB: db}

	err := tm.Execute(func(tx *gorm.DB) error {

		studentData := map[string]any{
			"StudentCode": *studentID,
			"FirstName":   *firstName,
			"LastName":    *lastName,
			// Add other fields as needed.
		}

		studentController := commonController.CreateStudentController(tx)
		if err := studentController.Update(*studentID, studentData); err != nil {
			return fmt.Errorf("failed to add student to common data: %w", err)
		}

		// Migrate the common student to HR.
		if err := controller.MigrateStudentsToHR(tx); err != nil {
			return fmt.Errorf("failed to migrate students to HR: %w", err)
		}

		// Update HR‑specific information.
		hrFacade := controller.NewHRFacade(tx)
		studentInfo, err := hrFacade.GetStudentById(*studentID)
		if err != nil {
			return fmt.Errorf("Error retrieving student with ID %s: %v\n", *studentID, err)
		}

		newStudent := hrModel.NewStudentInfoBuilder().
			WithFirstName(ifNotEmpty(*firstName, studentInfo.FirstName)).
			WithLastName(ifNotEmpty(*lastName, studentInfo.LastName)).
			WithGender(ifNotEmpty(*gender, studentInfo.Gender)).
			WithCitizenID(ifNotEmpty(*citizenID, studentInfo.CitizenID)).
			WithPhoneNumber(ifNotEmpty(*phoneNumber, studentInfo.PhoneNumber)).
			WithEmail(ifNotEmpty(*emailStudent, studentInfo.Email)).
			Build()

		if err := hrFacade.UpdateStudent(newStudent); err != nil {
			return fmt.Errorf("failed to update student info: %w", err)
		}

		return nil
	})

	if err != nil {
		return fmt.Errorf("Transaction failed: %v\n", err)
	}

	fmt.Println("Student updated successfully!")
	return nil
}
