package commands

import (
	"ModEd/hr/controller"
	"ModEd/hr/util"
	"fmt"

	"gorm.io/gorm"
)

// usage: go run hr/cli/HumanResourceCLI.go migrate
func (c *MigrateStudentsCommand) Execute(args []string, tx *gorm.DB) error {
	db := util.OpenDatabase(*util.DatabasePath)

	if err := controller.MigrateStudentsToHR(db); err != nil {
		return fmt.Errorf("Migration failed: %v\n", err)
	}

	fmt.Println("Migration completed successfully!")
	return nil
}
