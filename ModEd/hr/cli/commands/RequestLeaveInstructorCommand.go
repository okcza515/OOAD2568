package commands

import (
	"ModEd/hr/controller"
	hrModel "ModEd/hr/model"
	"ModEd/hr/util"
	"flag"
	"fmt"

	"gorm.io/gorm"
)

// usage: go run hr/cli/HumanResourceCLI.go request instructor leave -id="66050001" -type="sick" -reason="ไม่สบาย" -date="2025-04-20"
func requestLeaveInstructor(args []string, tx *gorm.DB) error {
	// fmt.Printf("asdjoasjdojaodjsaojdoasjodsjodasj")
	fs := flag.NewFlagSet("requestLeave", flag.ExitOnError)
	InstructorID := fs.String("id", "", "Instructor ID")
	leaveType := fs.String("type", "", "Type of leave (e.g. sick, personal)")
	reason := fs.String("reason", "", "Reason for leave")
	leaveDateStr := fs.String("date", "", "Leave date (YYYY-MM-DD)")
	fs.Parse(args)

	if err := util.ValidateRequiredFlags(fs, []string{"id", "type", "reason", "date"}); err != nil {
		fs.Usage()
		return fmt.Errorf("Validation error: %v\n", err)
	}

	db := util.OpenDatabase(*util.DatabasePath)
	hrFacade := controller.NewHRFacade(db)

	request := hrModel.NewRequestLeaveInstructorBuilder().
	WithInstructorID(*InstructorID).
	WithLeaveType(*leaveType).
	WithReason(*reason).
	WithLeaveDate(*leaveDateStr).
	Build()

	if err := hrFacade.SubmitLeaveInstructorRequest(request); err != nil {
		return fmt.Errorf("Failed to submit leave request: %v\n", err)
	}

	fmt.Println("Leave request submitted successfully.")
	return nil
}
