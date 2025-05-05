package commands

import (
	"gorm.io/gorm"
)

type ExportInstructorCommand struct{}

func (cmd *ExportInstructorCommand) Execute(args []string, tx *gorm.DB) error {
	// return handleExport(args, "export-instructor", tx)
	return nil
}
