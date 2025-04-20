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
