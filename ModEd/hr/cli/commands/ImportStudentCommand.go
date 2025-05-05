package commands

import (
	"ModEd/hr/controller"

	"gorm.io/gorm"
)

type ImportStudentCommand struct{}

// usage : go run hr/cli/HumanResourceCLI.go import student -path=<path>
// required field : path !!
func (cmd *ImportStudentCommand) Execute(args []string, tx *gorm.DB) error {
	controller := controller.NewStudentHRController(tx)
	return handleImport(args, "import-student", controller.ImportStudents, tx)
}
