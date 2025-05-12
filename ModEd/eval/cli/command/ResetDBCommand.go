package command

import (
	"ModEd/core"
	"ModEd/core/migration"
)

const (
	defaultDBPath = "../../data/ModEd.bin"
)

type ResetDBCommand struct{}

func (r *ResetDBCommand) Execute() error {
	err := migration.GetInstance().DropAllTables()
	if err != nil {
		return err
	}

	_, err = migration.GetInstance().
		SetPathDB(defaultDBPath).
		MigrateModule(core.MODULE_QUIZ).
		BuildDB()

	if err != nil {
		return err
	}

	return nil
}
