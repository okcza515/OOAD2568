package commands

import "gorm.io/gorm"

type RequestStudentResignCommand struct{}

func (cmd *RequestStudentResignCommand) Execute(args []string, tx *gorm.DB) error {
	return handleResignationRequest("student", args, tx)
}
