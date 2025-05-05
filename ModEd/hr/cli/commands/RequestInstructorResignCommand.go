package commands

import "gorm.io/gorm"

type RequestInstructorResignCommand struct{}

func (cmd *RequestInstructorResignCommand) Execute(args []string, tx *gorm.DB) error {
	return handleResignationRequest("instructor", args, tx)
}
