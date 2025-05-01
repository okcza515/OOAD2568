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

	validtor := util.NewValidationChain(fs)
	validtor.Field("path").Required()
	err := validtor.Validate()
	if err != nil {
		fs.Usage()
		return fmt.Errorf("validation error: %v", err)
	}

	if _, err := os.Stat(*filePath); errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("*** Error: File %s does not exist.\n", *filePath)
	}

	hrMapper, err := core.CreateMapper[model.InstructorInfo](*filePath)
	if err != nil {
		return fmt.Errorf("Failed to create HR mapper: %v\n", err)
	}

	instructors := hrMapper.Deserialize()

	tm := &util.TransactionManager{DB: tx}
	return tm.Execute(func(tx *gorm.DB) error {
		if err := controller.ImportInstructors(tx, instructors); err != nil {
			return err
		}
		fmt.Println("Instructors imported successfully!")
		return nil
	})
}
