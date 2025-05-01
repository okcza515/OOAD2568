package commands

import (
	"ModEd/hr/controller"
	"ModEd/hr/util"
	"flag"
	"fmt"

	"gorm.io/gorm"
)

func (c *ExportStudentsCommand) Execute(args []string, tx *gorm.DB) error {
	fs := flag.NewFlagSet("export", flag.ExitOnError)
	filePath := fs.String("path", "", "File path to export data")
	format := fs.String("format", "", "Export format (csv or json)")
	fs.Parse(args)

	validator := util.NewValidationChain(fs)
	validator.Field("path").Required()
	validator.Field("format").Required().AllowedValues([]string{"csv", "json"})
	err := validator.Validate()
	if err != nil {
		fs.Usage()
		return fmt.Errorf("validation error: %v", err)
	}

	// Delegate the export logic to the controller
	if err := controller.ExportStudents(tx, *filePath, *format); err != nil {
		return fmt.Errorf("failed to export students: %v", err)
	}

	fmt.Println("Student info exported successfully!")
	return nil
}
