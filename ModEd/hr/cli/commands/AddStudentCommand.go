package commands

import (
	commonController "ModEd/common/controller"
	commonModel "ModEd/common/model"
	"ModEd/hr/controller"
	hrModel "ModEd/hr/model"
	"ModEd/hr/util"
	"flag"
	"fmt"
	"os"

	"gorm.io/gorm"
)

func (c *AddStudentCommand) Run(args []string) {
	fs := flag.NewFlagSet("add", flag.ExitOnError)
	studentID := fs.String("id", "", "Student ID")
	firstName := fs.String("fname", "", "First Name")
	lastName := fs.String("lname", "", "Last Name")
	gender := fs.String("gender", "", "Gender")
	citizenID := fs.String("citizenID", "", "Citizen ID")
	phoneNumber := fs.String("phone", "", "Phone Number")
	fs.Parse(args)

	if err := util.ValidateRequiredFlags(fs, []string{"id", "fname", "lname"}); err != nil {
		fmt.Printf("Validation error: %v\n", err)
		fs.Usage()
		os.Exit(1)
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
		newStudent := hrModel.NewStudentInfoBuilder().
			WithStudentCode(*studentID).
			WithFirstName(*firstName).
			WithLastName(*lastName).
			WithGender(*gender).
			WithCitizenID(*citizenID).
			WithPhoneNumber(*phoneNumber).
			Build()

		if err := hrFacade.UpdateStudent(newStudent); err != nil {
			return fmt.Errorf("failed to update student info: %w", err)
		}

		return nil
	})

	if err != nil {
		fmt.Printf("Transaction failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Student added and HR info updated successfully!")
}
