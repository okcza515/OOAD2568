package commands

import (
	"ModEd/hr/controller"
	"ModEd/hr/util"

	"flag"
	"fmt"

	"gorm.io/gorm"
)

type ImportStudentCommand struct{}

// usage : go run hr/cli/HumanResourceCLI.go import student -path=<path>
// required field : path !!
func (cmd *ImportStudentCommand) Execute(args []string, tx *gorm.DB) error {
	fs := flag.NewFlagSet("import-student", flag.ExitOnError)
	filePath := fs.String("path", "", "Path to CSV or JSON for HR student info (only studentid and HR fields).")
	fs.Parse(args)

	validator := util.NewValidationChain(fs)
	validator.Field("path").Required()
	err := validator.Validate()
	if err != nil {
		fs.Usage()
		return fmt.Errorf("validation error: %v", err)
	}

	tm := &util.TransactionManager{DB: tx}
	return tm.Execute(func(tx *gorm.DB) error {
		studentController := controller.NewStudentHRController(tx)
		if err := studentController.ImportStudents(*filePath); err != nil {
			return err
		}
		fmt.Println("Students imported successfully!")
		return nil
	})
}
