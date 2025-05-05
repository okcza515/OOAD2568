package commands

import (
	"ModEd/hr/util"
	"flag"
	"fmt"

	"gorm.io/gorm"
)

type ControllerImportFunc func(filePath string) error

func handleImport(args []string, commandName string, controllerFunc ControllerImportFunc, tx *gorm.DB) error {
	fs := flag.NewFlagSet(commandName, flag.ExitOnError)
	path := fs.String("path", "", "Path to the file to import")

	if err := fs.Parse(args); err != nil {
		return fmt.Errorf("failed to parse flags: %v", err)
	}

	validator := util.NewValidationChain(fs)
	validator.Field("path").Required()
	if err := validator.Validate(); err != nil {
		fs.Usage()
		return fmt.Errorf("validation error: %v", err)
	}

	if err := controllerFunc(*path); err != nil {
		return fmt.Errorf("failed to import file: %v", err)
	}

	fmt.Println("Import successful")
	return nil
}
