package commands

import (
	"ModEd/hr/controller"

	"gorm.io/gorm"
)

// ReviewInstructorLeaveCommand processes the review for an instructor leave request.
// Usage: go run hr/cli/HumanResourceCLI.go review-instructor-leave -id=<requestID> -action=approve|reject [-reason="if reject"]
type ReviewInstructorRaiseCommand struct{}

func (cmd *ReviewInstructorRaiseCommand) Execute(args []string, tx *gorm.DB) error {
	controller := controller.CreateRaiseInstructorHRController(tx)
	return handleReviewCommand(
		args,
		tx,
		"review-instructor-raise",
		"Instructor raise request ID",
		controller.ReviewInstructorRaiseRequest,
		"instructor raise",
	)
}				