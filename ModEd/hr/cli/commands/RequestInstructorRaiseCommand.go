package commands

import (
	"gorm.io/gorm"
)

type RequestInstructorRaiseCommand struct{}

func (cmd *RequestInstructorRaiseCommand) Execute(args []string, tx *gorm.DB) error {
	return handleRaiseRequest("instructor", args, tx)
}