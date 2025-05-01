package commands

import (
	"ModEd/hr/controller"

	"ModEd/hr/util"
	"flag"
	"fmt"

	"gorm.io/gorm"
)

// usage: go run hr/cli/HumanResourceCLI.go request instructor raise -id="66050001" -amount=10000 -reason="ดีมาก"
func requestRaiseInstructor(args []string, tx *gorm.DB) error {
	fs := flag.NewFlagSet("requestRaise", flag.ExitOnError)
	id := fs.String("id", "", "Instructor ID")
	amount := fs.Int("amount", 0, "Raise amount")
	reason := fs.String("reason", "", "Reason for raise")

	if err := fs.Parse(args); err != nil {
		return fmt.Errorf("failed to parse flags: %v", err)
	}

	validator := util.NewValidationChain(fs)
	validator.Field("id").Required().Length(11).Regex(`^[0-9]{11}$`)
	validator.Field("amount").Required()
	validator.Field("reason").Required()
	err := validator.Validate()
	if err != nil {
		fs.Usage()
		return fmt.Errorf("validation error: %v", err)
	}

	raiseController := controller.NewRaiseHRController(tx) // Declare and initialize the controller
	if err := raiseController.SubmitRaiseRequest(*id, *amount, *reason); err != nil {
		return err
	}

	fmt.Println("Raise request submitted successfully.")
	return nil
}
