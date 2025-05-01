package commands

import (
	"flag"
	"fmt"

	"ModEd/hr/controller"
	"ModEd/hr/util"

	"gorm.io/gorm"
)

type DeleteStudentCommand struct{}

// usage : go run hr/cli/HumanResourceCLI.go delete -field="value"
// required field : id !!
func (cmd *DeleteStudentCommand) Execute(args []string, tx *gorm.DB) error {
	fs := flag.NewFlagSet("delete", flag.ExitOnError)
	studentID := fs.String("id", "", "Student ID to delete")
	fs.Parse(args)

	validator := util.NewValidationChain(fs)
	validator.Field("id").Required().Length(11).Regex(`^[0-9]{11}$`)
	err := validator.Validate()

	if err != nil {
		fs.Usage()
		return fmt.Errorf("validation error: %v", err)
	}

	err = controller.DeleteStudent(tx, *studentID)

	if err != nil {
		return fmt.Errorf("failed to delete student: %v", err)
	}

	fmt.Println("Student deleted successfully!")
	return nil
}
