package commands

import "gorm.io/gorm"

type RequestInstructorLeaveCommand struct{}

func (cmd *RequestInstructorLeaveCommand) Execute(args []string, tx *gorm.DB) error {
	return handleLeaveRequest("instructor", args, tx)
}
