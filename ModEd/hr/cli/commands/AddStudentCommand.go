package commands

import (
	commonController "ModEd/common/controller"
	commonModel "ModEd/common/model"
	"ModEd/hr/controller"
	hrModel "ModEd/hr/model"
	"ModEd/hr/util"
	"flag"
	"fmt"

	"gorm.io/gorm"
)

func (c *AddStudentCommand) Execute(args []string, tx *gorm.DB) error {
	fs := flag.NewFlagSet("add", flag.ExitOnError)
	studentID := fs.String("id", "", "Student ID")
	firstName := fs.String("fname", "", "First Name")
	lastName := fs.String("lname", "", "Last Name")
	gender := fs.String("gender", "", "Gender")
	citizenID := fs.String("citizenID", "", "Citizen ID")
	phoneNumber := fs.String("phone", "", "Phone Number")
	fs.Parse(args)

	if err := util.ValidateRequiredFlags(fs, []string{"id", "fname", "lname"}); err != nil {
		fs.Usage()
		return fmt.Errorf("validation error: %v", err)
	}

	db := util.OpenDatabase(*util.DatabasePath)

	// Create a TransactionManager instance.
	tm := &util.TransactionManager{DB: db}

	err := tm.Execute(func(tx *gorm.DB) error {
		// Create common student record.
		commonStudent := &commonModel.Student{
			StudentCode: *studentID,
			FirstName:   *firstName,
			LastName:    *lastName,
			// Populate additional fields if needed.
		}
		studentController := commonController.CreateStudentController(tx)
		if err := studentController.Create(commonStudent); err != nil {
			return fmt.Errorf("failed to add student to common data: %w", err)
		}

		// Migrate the common student to HR.
		if err := controller.MigrateStudentsToHR(tx); err != nil {
			return fmt.Errorf("failed to migrate students to HR: %w", err)
		}

		// Update HR‑specific information.
		hrFacade := controller.NewHRFacade(tx)
		builder := hrModel.NewStudentInfoBuilder()
		newStudent, err := builder.
			WithStudentCode(*studentID).
			WithFirstName(*firstName).
			WithLastName(*lastName).
			WithGender(*gender).
			WithCitizenID(*citizenID).
			WithPhoneNumber(*phoneNumber).
			Build()

		if err != nil {
			return fmt.Errorf("failed to build student info: %w", err)
		}

		if err := hrFacade.UpdateStudent(newStudent); err != nil {
			return fmt.Errorf("failed to update student info: %w", err)
		}

		return nil
	})

	if err != nil {
		return fmt.Errorf("transaction failed: %v", err)
	}

	fmt.Println("Student added and HR info updated successfully!")

	return nil
}
