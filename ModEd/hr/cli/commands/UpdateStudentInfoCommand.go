package commands

import (
	commonController "ModEd/common/controller"
	"ModEd/hr/controller"
	"ModEd/hr/model"
	"ModEd/hr/util"
	"flag"
	"fmt"

	"gorm.io/gorm"
)

// usage : go run hr/cli/HumanResourceCLI.go update student info -field="value"
// required field : id !!

func updateStudentInfo(args []string, tx *gorm.DB) error {
	fs := flag.NewFlagSet("update student", flag.ExitOnError)
	studentID := fs.String("id", "", "Student ID to update")
	firstName := fs.String("fname", "", "New first name")
	lastName := fs.String("lname", "", "New last name")
	gender := fs.String("gender", "", "New gender")
	citizenID := fs.String("citizenID", "", "New citizen ID")
	phoneNumber := fs.String("phone", "", "New phone number")
	email := fs.String("email", "", "New email")
	fs.Parse(args)

	if *studentID == "" {
		fs.Usage()
		return fmt.Errorf("student id is required")
	}

	tm := &util.TransactionManager{DB: tx}
	return tm.Execute(func(tx *gorm.DB) error {
		hrFacade := controller.NewHRFacade(tx)
		studentInfo, err := hrFacade.GetStudentById(*studentID)
		if err != nil {
			return fmt.Errorf("error retrieving student with ID %s: %v", *studentID, err)
		}

		// Create updated student info using non-empty flag values.
		builder := model.NewStudentInfoBuilder()
	
		updatedStudent, err := builder.
			WithStudentCode(*studentID).
			WithFirstName(ifNotEmpty(*firstName, studentInfo.FirstName)).
			WithLastName(ifNotEmpty(*lastName, studentInfo.LastName)).
			WithGender(ifNotEmpty(*gender, studentInfo.Gender)).
			WithCitizenID(ifNotEmpty(*citizenID, studentInfo.CitizenID)).
			WithPhoneNumber(ifNotEmpty(*phoneNumber, studentInfo.PhoneNumber)).
			WithEmail(ifNotEmpty(*email, studentInfo.Email)).
			Build()

		// Update common student data.
		studentData := map[string]any{
			"FirstName": updatedStudent.FirstName,
			"LastName":  updatedStudent.LastName,
			// add additional fields as needed.
		}
		studentController := commonController.CreateStudentController(tx)
		if err := studentController.Update(*studentID, studentData); err != nil {
			return fmt.Errorf("failed to update common student data: %v", err)
		}

		// Migrate and update HR-specific data.
		if err := controller.MigrateStudentsToHR(tx); err != nil {
			return fmt.Errorf("failed to migrate student to HR module: %v", err)
		}

		if err := hrFacade.UpdateStudent(updatedStudent); err != nil {
			return fmt.Errorf("failed to update student HR info: %v", err)
		}
		fmt.Println("Student updated successfully!")
		return nil
	})
}

// ifNotEmpty returns newValue if not empty, otherwise fallback.
func ifNotEmpty(newValue, fallback string) string {
	if newValue != "" {
		return newValue
	}
	return fallback
}
