package commands

import (
	"ModEd/hr/controller"

	"gorm.io/gorm"
)

// ReviewInstructorLeaveCommand processes the review for an instructor leave request.
// Usage: go run hr/cli/HumanResourceCLI.go review-instructor-leave -id=<requestID> -action=approve|reject [-reason="if reject"]
type ReviewInstructorLeaveCommand struct{}

func (cmd *ReviewInstructorLeaveCommand) Execute(args []string, tx *gorm.DB) error {
	// TODO: Add controller method
	controller := controller.CreateLeaveInstructorHRController(tx)
	return handleReviewCommand(
		args,
		tx,
		"review-instructor-leave",
		"Instructor leave request ID",
		controller.ReviewInstructorLeaveRequest,
		"instructor leave",
	)
}
