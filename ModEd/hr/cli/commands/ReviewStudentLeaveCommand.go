package commands

import (
	"gorm.io/gorm"
)

type ReviewStudentLeaveCommand struct{}

func (cmd *ReviewStudentLeaveCommand) Execute(args []string, tx *gorm.DB) error {
	// TODO: Add controller method
	return handleReviewCommand(
		args,
		tx,
		"review-student-leave",
		"Student leave request ID",
		controller.ReviewStudentLeaveRequest,
		"student leave",
	)
}
