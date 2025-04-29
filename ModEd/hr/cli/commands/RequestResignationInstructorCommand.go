package commands

import (
	"ModEd/hr/controller"
	"ModEd/hr/util"
	"flag"
	"fmt"

	"gorm.io/gorm"
)

// usage: go run hr/cli/HumanResourceCLI.go request instructor resign -id="66050001" -reason="เหนื่อยมาก"
func requestResignationInstructor(args []string, tx *gorm.DB) error {
	fs := flag.NewFlagSet("requestResignation", flag.ExitOnError)
	id := fs.String("id", "", "ID")
	reason := fs.String("reason", "", "Reason for resignation")

	if err := fs.Parse(args); err != nil {
		return fmt.Errorf("failed to parse flags: %v", err)
	}

	err := util.NewValidationChain(fs).
		Required("reason").
		Validate()
	if err != nil {
		fs.Usage()
		return fmt.Errorf("validation error: %v", err)
	}

	hrFacade := controller.NewHRFacade(tx)
	if err := hrFacade.SubmitResignationInstructor(tx, *id, *reason); err != nil {
		return fmt.Errorf("failed to submit resignation request: %v", err)
	}

	fmt.Println("Resignation request submitted successfully.")
	return nil
}
