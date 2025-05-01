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

	validator := util.NewValidationChain(fs)
	validator.Field("answer").Required().AllowedValues([]string{"approve", "reject"})
	validator.Field("reason").Required()
	err := validator.Validate()
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

	hrFacade := controller.NewHRFacade(tx)
	if err := hrFacade.UpdateResignationStudentStatus(*id, status, *reason); err != nil {
		return fmt.Errorf("failed to update resignation status: %v", err)
	}

	fmt.Printf("Resignation marked as '%s' for id %s\n", status, *id)
	return nil
}
