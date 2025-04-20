package commands

import (
	commonController "ModEd/common/controller"
	"ModEd/hr/controller"
	hrUtil "ModEd/hr/util"
	"flag"
	"fmt"
	"os"

	"gorm.io/gorm"
)

// usage : go run hr/cli/HumanResourceCLI.go delete -field="value"
// required field : id !!"
func (c *DeleteStudentCommand) Run(args []string) {
	fs := flag.NewFlagSet("delete", flag.ExitOnError)
	studentID := fs.String("id", "", "Student ID to delete")
	fs.Parse(args)

	if err := hrUtil.ValidateRequiredFlags(fs, []string{"id"}); err != nil {
		fmt.Printf("Validation error: %v\n", err)
		fs.Usage()
		os.Exit(1)
	}

	db := hrUtil.OpenDatabase(*hrUtil.DatabasePath)

	// Create a TransactionManager instance.
	tm := &hrUtil.TransactionManager{DB: db}

	err := tm.Execute(func(tx *gorm.DB) error {
		// Delete student from common data.
		studentController := commonController.CreateStudentController(tx)
		if err := studentController.DeleteByStudentId(*studentID); err != nil {
			return fmt.Errorf("failed to delete student from common data: %w", err)
		}

		// Migrate the common student to HR.
		if err := controller.MigrateStudentsToHR(tx); err != nil {
			return fmt.Errorf("failed to migrate students to HR: %w", err)
		}

		// Delete HR-specific information.
		hrFacade := controller.NewHRFacade(tx)
		if err := hrFacade.DeleteStudent(*studentID); err != nil {
			return fmt.Errorf("failed to delete student from HR data: %w", err)
		}

		return nil
	})

	if err != nil {
		fmt.Printf("Failed to delete student: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Student deleted successfully!")
}
