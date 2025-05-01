package commands

import (
	"gorm.io/gorm"
)

type RequestStudentLeaveCommand struct{}

func (cmd *RequestStudentLeaveCommand) Execute(args []string, tx *gorm.DB) error {
	return handleLeaveRequest("student", args, tx)
}
