package commands

import (
	"ModEd/hr/controller"
	model "ModEd/hr/model"

	"ModEd/hr/util"
	"flag"
	"fmt"

	"gorm.io/gorm"
)

// usage: go run hr/cli/HumanResourceCLI.go request instructor raise -id="66050001" -amount=10000 -reason="ดีมาก"
func requestRaiseInstructor(args []string, tx *gorm.DB) error {
	fs := flag.NewFlagSet("requestRaise", flag.ExitOnError)
	ID := fs.String("id", "", "Instructor ID")
	amount := fs.Int("amount", 0, "Raise amount")
	reason := fs.String("reason", "", "Reason for raise")

	if err := fs.Parse(args); err != nil {
		return fmt.Errorf("failed to parse flags: %v", err)
	}

	if err := util.ValidateRequiredFlags(fs, []string{"id", "amount", "reason"}); err != nil {
		fs.Usage()
		return fmt.Errorf("validation error: %v", err)
	}

	tm := &util.TransactionManager{DB: tx}
	return tm.Execute(func(tx *gorm.DB) error {
		hrFacade := controller.NewHRFacade(tx)

		request := model.NewRequestRaiseBuilder().
			WithInstructorCode(*ID).
			WithTargetSalary(int(*amount)).
			WithReason(*reason).
			Build()

		if err := hrFacade.SubmitRaiseInstructorRequest(request); err != nil {
			return fmt.Errorf("failed to submit raise request: %v", err)
		}

		fmt.Println("Raise request submitted successfully.")
		return nil
	})
}
