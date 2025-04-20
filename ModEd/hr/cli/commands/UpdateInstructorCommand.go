package commands

import (
	"flag"
	"fmt"
	"strings"

	"ModEd/hr/controller"
	"ModEd/hr/model"
	"ModEd/hr/util"

	"gorm.io/gorm"
)

func (c *UpdateInstructorCommand) Execute(args []string, tx *gorm.DB) error {
	fs := flag.NewFlagSet("update", flag.ExitOnError)
	instructorID := fs.String("id", "", "Instructor ID")
	field := fs.String("field", "", "Field to update (position, department, etc.)")
	value := fs.String("value", "", "New value for the specified field")
	fs.Parse(args)

	if *instructorID == "" || *field == "" || *value == "" {
		fs.Usage()
		return fmt.Errorf("Instructor ID, field, and value are all required.")
	}

	db := util.OpenDatabase(*util.DatabasePath)
	tm := &util.TransactionManager{DB: db}
	err := tm.Execute(func(tx *gorm.DB) error {
		hrFacade := controller.NewHRFacade(tx)
		instructorInfo, err := hrFacade.GetInstructorById(*instructorID)
		if err != nil {
			return fmt.Errorf("error retrieving instructor: %v", err)
		}

		// Determine which field to update based on the "field" flag.
		switch strings.ToLower(*field) {
		case "position", "academicposition", "academic_position":
			// Convert string to AcademicPosition type.
			parsedPos, err := model.ParseAcademicPosition(*value)
			if err != nil {
				return fmt.Errorf("invalid academic position: %v", err)
			}
			instructorInfo.AcademicPosition = parsedPos
		case "department":
			// Assume InstructorInfo has a Department field.
			// instructorInfo.Department = *value
		// Add more cases as needed for additional fields.
		default:
			return fmt.Errorf("unknown field: %s", *field)
		}

		if err := hrFacade.UpdateInstructor(instructorInfo); err != nil {
			return fmt.Errorf("error updating instructor: %v", err)
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("Transaction failed: %v\n", err)
	}
	fmt.Printf("Instructor %s updated successfully: %s set to %s\n", *instructorID, *field, *value)
	return nil
}
