package commands

import (
	"ModEd/core"
	"ModEd/core/migration"
	"fmt"

	"gorm.io/gorm"
)

type MigrateCommand struct{}

// usage: go run hr/cli/HumanResourceCLI.go migrate
func (cmd *MigrateCommand) Execute(args []string, tx *gorm.DB) error {
	mm := migration.GetInstance()
	mm.MigrateModule(core.MODULE_HR)
	_, err := mm.BuildDB()
	if err != nil {
		return err
	}
	fmt.Println("Database migration completed successfully.")
	return nil
}
