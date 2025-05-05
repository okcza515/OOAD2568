package commands

import (
	"ModEd/hr/controller"

	"gorm.io/gorm"
)

// ReviewInstructorResignCommand processes the review for an instructor resignation request.
// Usage: go run hr/cli/HumanResourceCLI.go review-instructor-resign -id=<requestID> -action=approve|reject [-reason="if reject"]
type ReviewInstructorResignCommand struct{}

func (cmd *ReviewInstructorResignCommand) Execute(args []string, tx *gorm.DB) error {
	controller := controller.NewResignationInstructorHRController(tx)
	return handleReviewCommand(
		args,
		tx,
		"review-instructor-resign",
		"Instructor resign request ID",
		controller.ReviewInstructorResignRequest,
		"instructor resign",
	)
}
