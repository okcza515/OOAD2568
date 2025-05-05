package commands

import (
	"ModEd/hr/controller"
	"ModEd/hr/util"
	"flag"
	"fmt"

	"gorm.io/gorm"
)

type RequestInstructorRaiseCommand struct{}

// usage: go run hr/cli/HumanResourceCLI.go request instructor raise -id="66050001" -amount=10000 -reason="ดีมาก"
func (cmd *RequestInstructorRaiseCommand) Execute(args []string, tx *gorm.DB) error {
	fs := flag.NewFlagSet("requestRaise", flag.ExitOnError)
	id := fs.String("id", "", "Instructor ID")
	amount := fs.Int("amount", 0, "Raise amount")
	reason := fs.String("reason", "", "Reason for raise")

	if err := fs.Parse(args); err != nil {
		return fmt.Errorf("failed to parse flags: %v", err)
	}

	validator := util.NewValidationChain(fs)
	// Instructor Code in demo data is different from real instructor ID in LEB2.
	validator.Field("id").Required().IsInstructorID()
	validator.Field("amount").Required()
	validator.Field("reason").Required()
	err := validator.Validate()
	if err != nil {
		fs.Usage()
		return fmt.Errorf("validation error: %v", err)
	}

	controller := controller.NewRaiseHRController(tx)

	err = controller.SubmitRaiseRequest(*id, *amount, *reason)
	if err != nil {
		return fmt.Errorf("failed to submit raise request: %v", err)
	}

	fmt.Println("Raise request submitted successfully.")
	return nil
}
