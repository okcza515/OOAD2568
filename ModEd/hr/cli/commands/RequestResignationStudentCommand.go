package commands

import (
	"ModEd/hr/controller"
	"ModEd/hr/util"
	"flag"
	"fmt"

	"gorm.io/gorm"
)

// usage: go run hr/cli/HumanResourceCLI.go request student resign -id="66050001" -reason="อยากลาออกโว้ย"
func requestResignationStudent(args []string, tx *gorm.DB) error {
	fs := flag.NewFlagSet("requestResignation", flag.ExitOnError)
	id := fs.String("id", "", "Student ID")
	reason := fs.String("reason", "", "Reason for resignation")

	if err := fs.Parse(args); err != nil {
		return fmt.Errorf("failed to parse flags: %v", err)
	}

	if err := util.NewValidationChain(fs).
		Required("id").
		Length("id", 11).
		Required("reason").
		Validate(); err != nil {
		fs.Usage()
		return fmt.Errorf("validation error: %v", err)
	}

	hrFacade := controller.NewHRFacade(tx)
	if err := hrFacade.SubmitResignationStudent(tx, *id, *reason); err != nil {
		return fmt.Errorf("failed to submit resignation request: %v", err)
	}
	fmt.Println("Resignation request submitted successfully.")
	return nil
}
