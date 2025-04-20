package commands

import (
	commonController "ModEd/common/controller"
	"ModEd/hr/controller"
	"ModEd/hr/model"
	"ModEd/hr/util"
	"flag"
	"fmt"
	"strings"

	"gorm.io/gorm"
)

func (c *UpdateCommand) Execute(args []string, tx *gorm.DB) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: update {student|instructor} [options]")
	}

	target := strings.ToLower(args[0])
	switch target {
	case "student":
		return updateStudent(args[1:], tx)
	case "instructor":
		return updateInstructor(args[1:], tx)
	default:
		return fmt.Errorf("unknown update target: %s", target)
	}
}

func updateStudent(args []string, tx *gorm.DB) error {
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
		updatedStudent := model.NewStudentInfoBuilder().
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

func updateInstructor(args []string, tx *gorm.DB) error {
	fs := flag.NewFlagSet("update instructor", flag.ExitOnError)
	instructorID := fs.String("id", "", "Instructor ID to update")
	field := fs.String("field", "", "Field to update (e.g., position, department)")
	value := fs.String("value", "", "New value for the specified field")
	fs.Parse(args)

	if *instructorID == "" || *field == "" || *value == "" {
		fs.Usage()
		return fmt.Errorf("instructor id, field, and value are required")
	}

	tm := &util.TransactionManager{DB: tx}
	return tm.Execute(func(tx *gorm.DB) error {
		hrFacade := controller.NewHRFacade(tx)
		instructorInfo, err := hrFacade.GetInstructorById(*instructorID)
		if err != nil {
			return fmt.Errorf("error retrieving instructor with ID %s: %v", *instructorID, err)
		}

		switch strings.ToLower(*field) {
		case "position", "academicposition", "academic_position":
			parsedPos, err := model.ParseAcademicPosition(*value)
			if err != nil {
				return fmt.Errorf("invalid academic position: %v", err)
			}
			instructorInfo.AcademicPosition = parsedPos
		case "department":
			// Assuming InstructorInfo has a Department field.
			// instructorInfo.Department = *value
		default:
			return fmt.Errorf("unknown field for instructor update: %s", *field)
		}

		if err := hrFacade.UpdateInstructor(instructorInfo); err != nil {
			return fmt.Errorf("error updating instructor: %v", err)
		}
		fmt.Println("Instructor updated successfully!")
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
