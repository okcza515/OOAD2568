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

// usage: go run hr/cli/HumanResourceCLI.go requestLeaveStudent -id="66050001" -type="sick" -reason="ไม่สบาย" -date="2025-04-20"
func RequestLeaveStudent(args []string, tx *gorm.DB) error {
	// fmt.Printf("asdjoasjdojaodjsaojdoasjodsjodasj")
	fs := flag.NewFlagSet("requestLeave", flag.ExitOnError)
	studentID := fs.String("id", "", "Student ID")
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
	request := hrModel.NewRequestLeaveStudentBuilder().
		WithStudentID(*studentID).
		WithLeaveType(*leaveType).
		WithReason(*reason).
		WithLeaveDate(leaveDate).
		Build()

	if err := hrFacade.SubmitLeaveStudentRequest(request); err != nil {
		return fmt.Errorf("Failed to submit leave request: %v\n", err)
	}

	fmt.Println("Leave request submitted successfully.")
	return nil
}
