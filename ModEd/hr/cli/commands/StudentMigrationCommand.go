package commands

import (
	"ModEd/hr/controller"
	"fmt"

	"gorm.io/gorm"
)

type MigrateStudentsCommand struct{}

// usage: go run hr/cli/HumanResourceCLI.go migrate
func (cmd *MigrateStudentsCommand) Execute(args []string, tx *gorm.DB) error {
	if err := controller.MigrateStudentsToHR(tx); err != nil {
		return fmt.Errorf("migration failed: %v", err)
	}

	fmt.Println("migration completed successfully!")
	return nil
}
