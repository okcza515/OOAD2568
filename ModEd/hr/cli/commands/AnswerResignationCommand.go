package commands

import (
	"ModEd/hr/controller"
	"ModEd/hr/util"
	"flag"
	"fmt"
	"os"
)

// usage : go run hr/cli/HumanResourceCLI.go answerResignation -id=3 -answer=approve
func (c *AnswerResignationCommand) Run(args []string) {
	fs := flag.NewFlagSet("answerResignation", flag.ExitOnError)
	id := fs.String("id", "", "id")
	answer := fs.String("answer", "", "approve or reject")
	reason := fs.String("reason", "", "Reason if rejected (optional)")
	fs.Parse(args)

	if err := util.ValidateRequiredFlags(fs, []string{"id", "answer"}); err != nil {
		fmt.Printf("Validation error: %v\n", err)
		fs.Usage()
		os.Exit(1)
	}

	var status string
	switch *answer {
	case "approve":
		status = "Approved"
	case "reject":
		status = "Rejected"
	default:
		fmt.Println("Error: --answer must be either 'approve' or 'reject'")
		os.Exit(1)
	}

	db := util.OpenDatabase(*util.DatabasePath)
	hrFacade := controller.NewHRFacade(db)

	err := hrFacade.UpdateResignationStatus(*id, status, *reason)
	if err != nil {
		fmt.Printf("Error updating resignation request: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Resignation marked as '%s' for id %s\n", status, *id)
}
