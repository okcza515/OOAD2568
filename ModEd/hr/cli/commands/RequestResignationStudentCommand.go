package commands

import (
	"ModEd/hr/controller"
	"ModEd/hr/model"
	"ModEd/hr/util"
	"flag"
	"fmt"

	"gorm.io/gorm"
)

// usage: go run hr/cli/HumanResourceCLI.go request student resign -id="66050001" -reason="อยากลาออกโว้ย"
func requestResignationStudent(args []string, tx *gorm.DB) error {
	fs := flag.NewFlagSet("requestResignation", flag.ExitOnError)
	ID := fs.String("id", "", "ID")
	reason := fs.String("reason", "", "Reason for resignation")

	if err := fs.Parse(args); err != nil {
		return fmt.Errorf("failed to parse flags: %v", err)
	}

	if err := util.ValidateRequiredFlags(fs, []string{"id", "reason"}); err != nil {
		fs.Usage()
		return fmt.Errorf("validation error: %v", err)
	}

	tm := &util.TransactionManager{DB: tx}
	return tm.Execute(func(tx *gorm.DB) error {
		hrFacade := controller.NewHRFacade(tx)

		builder := model.NewRequestResignationBuilder(true)
		req, err := builder.
			WithID(*ID).
			WithReason(*reason).
			Build()

		if err != nil {
			return fmt.Errorf("failed to build resignation request: %v", err)
		}

		if err := hrFacade.SubmitResignationStudentRequest(req.(*model.RequestResignationStudent)); err != nil {
			return fmt.Errorf("failed to submit resignation request: %v", err)
		}

		fmt.Println("Resignation request submitted successfully.")
		return nil
	})
}
