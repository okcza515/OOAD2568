package commands

import (
	"ModEd/hr/controller"
	"ModEd/hr/util"
	"flag"
	"fmt"

	"gorm.io/gorm"
)

// usage : go run hr/cli/HumanResourceCLI.go answerResignation -id=3 -answer=approve
func (c *AnswerResignationCommand) Execute(args []string, tx *gorm.DB) error {
	fs := flag.NewFlagSet("answerResignation", flag.ExitOnError)
	id := fs.String("id", "", "id")
	answer := fs.String("answer", "", "approve or reject")
	reason := fs.String("reason", "", "Reason if rejected (optional)")
	fs.Parse(args)

	err := util.NewValidationChain(fs).
		Required("answer").
		Required("reason").
		Validate()
	if err != nil {
		fs.Usage()
		return fmt.Errorf("validation error: %v", err)
	}

	var status string
	switch *answer {
	case "approve":
		status = "Approved"
	case "reject":
		status = "Rejected"
	default:
		return fmt.Errorf("invalid answer: --answer must be either 'approve' or 'reject'")
	}

	err = controller.HandleResignationStatus(tx, *id, status, *reason)
	if err != nil {
		return fmt.Errorf("failed to process resignation: %w", err)
	}

	fmt.Printf("Resignation marked as '%s' for id %s\n", status, *id)
	return nil
}
