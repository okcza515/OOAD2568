package commands

import (
	"ModEd/hr/controller"

	"gorm.io/gorm"
)

type ReviewStudentLeaveCommand struct{}

func (cmd *ReviewStudentLeaveCommand) Execute(args []string, tx *gorm.DB) error {
	controller := controller.CreateLeaveStudentHRController(tx)
	return handleReviewCommand(
		args,
		tx,
		"review-student-leave",
		"Student leave request ID",
		controller.ReviewStudentLeaveRequest,
		"student leave",
	)
}
