package commands

import (
	"ModEd/hr/controller"

	"gorm.io/gorm"
)

// ReviewStudentResignCommand processes the review for a student resignation request.
// Usage: go run hr/cli/HumanResourceCLI.go review-student-resign -id=<requestID> -action=approve|reject [-reason="if reject"]
type ReviewStudentResignCommand struct{}

func (cmd *ReviewStudentResignCommand) Execute(args []string, tx *gorm.DB) error {
	controller := controller.CreateResignationStudentHRController(tx)
	return handleReviewCommand(
		args,
		tx,
		"review-student-resign",
		"Student resign request ID",
		controller.ReviewStudentResignRequest,
		"student resign",
	)
}
