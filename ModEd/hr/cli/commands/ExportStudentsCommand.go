package commands

import (
	"gorm.io/gorm"
)

type ExportStudentsCommand struct{}

func (cmd *ExportStudentsCommand) Execute(args []string, tx *gorm.DB) error {
	// return handleExport(args, "export-students", tx)
	return nil
}
