package commands

import (
	"ModEd/hr/controller"
	hrModel "ModEd/hr/model"
	"ModEd/hr/util"
	"flag"
	"fmt"
	"time"

	"gorm.io/gorm"
)

// usage: go run hr/cli/HumanResourceCLI.go requestLeaveInstructor -id="66050001" -type="sick" -reason="ไม่สบาย" -date="2025-04-20"
func RequestLeaveInstructor(args []string, tx *gorm.DB) error {
	// fmt.Printf("asdjoasjdojaodjsaojdoasjodsjodasj")
	fs := flag.NewFlagSet("requestLeave", flag.ExitOnError)
	InstructorID := fs.String("id", "", "Instructor ID")
	leaveType := fs.String("type", "", "Type of leave (e.g. sick, personal)")
	reason := fs.String("reason", "", "Reason for leave")
	leaveDateStr := fs.String("date", "", "Leave date (YYYY-MM-DD)")
	fs.Parse(args)

	// fmt.Printf("Student ID: %s\n", *studentID)

	if err := util.ValidateRequiredFlags(fs, []string{"id", "type", "reason", "date"}); err != nil {
		fs.Usage()
		return fmt.Errorf("Validation error: %v\n", err)
	}

	// แปลง string -> time.Time
	leaveDate, err := time.Parse("2006-01-02", *leaveDateStr)
	if err != nil {
		return fmt.Errorf("Invalid date format: %v\n", err)
	}

	// เปิด database และเตรียม facade
	db := util.OpenDatabase(*util.DatabasePath)
	hrFacade := controller.NewHRFacade(db)

	// สร้างคำขอลา
	request := hrModel.NewRequestLeaveInstructorBuilder().
	WithInstructorID(*InstructorID).
	WithLeaveType(*leaveType).
	WithReason(*reason).
	WithLeaveDate(leaveDate).
	Build()

	if err := hrFacade.SubmitLeaveInstructorRequest(request); err != nil {
		return fmt.Errorf("Failed to submit leave request: %v\n", err)
	}

	fmt.Println("Leave request submitted successfully.")
	return nil
}
