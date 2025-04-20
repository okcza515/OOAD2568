package commands

import (
	"ModEd/hr/controller"
	hrModel "ModEd/hr/model"
	"ModEd/hr/util"
	"flag"
	"fmt"
	"os"
)


// usage: go run hr/cli/HumanResourceCLI.go requestResignation -id="66050001" -reason="ย้ายคณะ"
func (c *RequestResignationCommand) Run(args []string) {
	fs := flag.NewFlagSet("requestResignation", flag.ExitOnError)
	studentID := fs.String("id", "", "Student ID")
	reason := fs.String("reason", "", "Reason for resignation (optional)")
	fs.Parse(args)

	if err := util.ValidateRequiredFlags(fs, []string{"id","reason"}); err != nil {
		fmt.Printf("Validation error: %v\n", err)
		fs.Usage()
		os.Exit(1)
	}

	db := util.OpenDatabase(*util.DatabasePath)
	hrFacade := controller.NewHRFacade(db)

	// สร้าง resignation request object
	request := hrModel.NewRequestResignationBuilder().
		WithStudentID(*studentID).
		WithReason(*reason).
		Build()

	if err := hrFacade.SubmitResignationRequest(request); err != nil {
		fmt.Printf("Failed to submit resignation request: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Resignation request submitted successfully.")
}
