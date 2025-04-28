package commands

import (
	"ModEd/core"
	"ModEd/hr/controller"
	"ModEd/hr/model"
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

	hrMapper, err := core.CreateMapper[model.StudentInfo](*filePath)
	if err != nil {
		return fmt.Errorf("failed to create HR mapper: %v", err)
	}

	hrRecords := hrMapper.Deserialize()
	hrRecordsMap := make(map[string]model.StudentInfo)
	for _, hrRec := range hrRecords {
		if _, exists := hrRecordsMap[hrRec.StudentCode]; exists {
			return fmt.Errorf("duplicate student code found: %s", hrRec.StudentCode)
		}
		hrRecordsMap[hrRec.StudentCode] = *hrRec
	}

	// db := util.OpenDatabase(*util.DatabasePath)

	tm := &util.TransactionManager{DB: tx} // use passed transaction connection
	return tm.Execute(func(tx *gorm.DB) error {
		hrFacade := controller.NewHRFacade(tx) // pass tx instead of new db
		if err := hrFacade.ImportStudents(tx, hrRecordsMap); err != nil {
			return err
		}
		fmt.Println("Students imported successfully!")
		return nil
	})
}
