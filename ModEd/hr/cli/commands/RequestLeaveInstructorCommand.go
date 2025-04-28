package commands

import (
	"ModEd/hr/controller"
	"ModEd/hr/util"
	"flag"
	"fmt"

	"gorm.io/gorm"
)

// usage: go run hr/cli/HumanResourceCLI.go request instructor leave -id="66050001" -type="sick" -reason="ไม่สบาย" -date="2025-04-20"
func requestLeaveInstructor(args []string, tx *gorm.DB) error {
	fs := flag.NewFlagSet("requestLeave", flag.ExitOnError)
	instructorID := fs.String("id", "", "Instructor ID")
	leaveType := fs.String("type", "", "Type of leave (e.g. sick, personal)")
	reason := fs.String("reason", "", "Reason for leave")
	leaveDateStr := fs.String("date", "", "Leave date (YYYY-MM-DD)")
	fs.Parse(args)

	err := util.NewValidationChain(fs).
		Required("id").
		Length("id", 11).
		Required("type").
		Required("reason").
		Required("date").
		Validate()
	if err != nil {
		fs.Usage()
		return fmt.Errorf("validation error: %v", err)
	}

	hrFacade := controller.NewHRFacade(tx)
	if err := hrFacade.SubmitInstructorLeaveRequest(tx,*instructorID, *leaveType, *reason, *leaveDateStr); err != nil {
		return err
	}

	fmt.Println("Leave request submitted successfully.")
	return nil
}
