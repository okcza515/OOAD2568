package handler

import (
	"ModEd/core"
	"ModEd/core/migration"
	"fmt"
)

type MigrationHandler struct{}

func NewMigrationHandlerStrategy() *MigrationHandler {
	return &MigrationHandler{}
}

func (handler MigrationHandler) Execute() error {
	mm := migration.GetInstance()
	mm.MigrateModule(core.MODULE_HR)
	_, err := mm.BuildDB()
	if err != nil {
		return err
	}
	fmt.Println("Database migration completed successfully.")
	return nil
}
