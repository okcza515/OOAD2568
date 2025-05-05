package commands

import (
	"ModEd/hr/controller"

	"gorm.io/gorm"
)

type ImportInstructorCommand struct{}

func (cmd *ImportInstructorCommand) Execute(args []string, tx *gorm.DB) error {
	controller := controller.NewInstructorHRController(tx)
	return handleImport(args, "import-instructor", controller.ImportInstructors, tx)
}
