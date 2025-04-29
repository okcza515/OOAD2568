package commands

import (
	"flag"
	"fmt"

	"ModEd/hr/controller"
	"ModEd/hr/util"

	"gorm.io/gorm"
)

// usage: go run hr/cli/HumanResourceCLI.go update instructor info -id="66050001" -field="position" -value="Professor"
func updateInstructorInfo(args []string, tx *gorm.DB) error {
	fs := flag.NewFlagSet("update instructor", flag.ExitOnError)
	instructorID := fs.String("id", "", "Instructor ID to update")
	field := fs.String("field", "", "Field to update (e.g., position, department)")
	value := fs.String("value", "", "New value for the specified field")
	fs.Parse(args)

	err := util.NewValidationChain(fs).
		Required("id").
		Length("id", 11).
		Required("field").
		Required("value").
		Validate()
	if err != nil {
		fs.Usage()
		return fmt.Errorf("validation error: %v", err)
	}

	if err := controller.UpdateInstructorInfo(tx, *instructorID, *field, *value); err != nil {
		return fmt.Errorf("failed to update instructor info: %v", err)
	}

	fmt.Println("Instructor updated successfully!")
	return nil
}
