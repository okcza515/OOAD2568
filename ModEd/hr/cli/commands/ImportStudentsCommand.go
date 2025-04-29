package commands

import (
	"ModEd/hr/controller"
	"ModEd/hr/util"

	"flag"
	"fmt"

	"gorm.io/gorm"
)

// usage : go run hr/cli/HumanResourceCLI.go import student -path=<path>
// required field : path !!

func importStudents(args []string, tx *gorm.DB) error {
	fs := flag.NewFlagSet("import", flag.ExitOnError)
	filePath := fs.String("path", "", "Path to CSV or JSON for HR student info (only studentid and HR fields).")
	fs.Parse(args)

	err := util.NewValidationChain(fs).
		Required("path").
		Validate()
	if err != nil {
		fs.Usage()
		return fmt.Errorf("validation error: %v", err)
	}

	tm := &util.TransactionManager{DB: tx} // use passed transaction connection
	return tm.Execute(func(tx *gorm.DB) error {
		hrFacade := controller.NewHRFacade(tx) // pass tx instead of new db
		if err := hrFacade.ImportStudents(tx, *filePath); err != nil {
			return err
		}
		fmt.Println("Students imported successfully!")
		return nil
	})
}
