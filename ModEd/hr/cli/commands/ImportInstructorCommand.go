package commands

import (
	"ModEd/core"
	"ModEd/hr/controller"
	"ModEd/hr/model"
	"ModEd/hr/util"
	"errors"
	"flag"
	"fmt"
	"os"

	"gorm.io/gorm"
)

func importInstructor(args []string, tx *gorm.DB) error {
	fs := flag.NewFlagSet("import", flag.ExitOnError)
	filePath := fs.String("path", "", "Path to CSV or JSON for HR instructor info (only instructorid and HR fields).")
	fs.Parse(args)

	if err := util.ValidateRequiredFlags(fs, []string{"path"}); err != nil {
		fs.Usage()
		return fmt.Errorf("Validation error: %v\n", err)
	}

	if _, err := os.Stat(*filePath); errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("*** Error: File %s does not exist.\n", *filePath)
	}

	// Use CreateMapper to get the appropriate mapper
	hrMapper, err := core.CreateMapper[model.InstructorInfo](*filePath)
	if err != nil {
		return fmt.Errorf("Failed to create HR mapper: %v\n", err)
	}

	// Deserialize the data
	instructors := hrMapper.Deserialize()

	// Validate and insert each instructor
	hrFacade := controller.NewHRFacade(tx)
	for _, instructor := range instructors {
		if instructor.ID == 0 || instructor.FirstName == "" {
			return fmt.Errorf("Invalid instructor data: %+v\n", instructor)
		}

		// Insert instructor into the database
		if err := hrFacade.InsertInstructor(instructor); err != nil {
			return fmt.Errorf("Failed to insert instructor %d: %v\n", instructor.ID, err)
		}
	}

	fmt.Println("Instructors imported successfully!")
	return nil
}
