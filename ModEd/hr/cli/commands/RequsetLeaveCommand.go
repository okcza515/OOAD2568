package commands

import (
	"ModEd/hr/controller"
	hrModel "ModEd/hr/model"
	"ModEd/hr/util"
	"flag"
	"fmt"
	"os"
	"time"
)

// usage: go run hr/cli/HumanResourceCLI.go requestLeave -id="66050001" -type="sick" -reason="ไม่สบาย" -date="2025-04-20"
func (c *RequsetLeaveCommand) Run(args []string) {
	fs := flag.NewFlagSet("requestLeave", flag.ExitOnError)
	studentID := fs.String("id", "", "Student ID")
	leaveType := fs.String("type", "", "Type of leave (e.g. sick, personal)")
	reason := fs.String("reason", "", "Reason for leave")
	leaveDateStr := fs.String("date", "", "Leave date (YYYY-MM-DD)")
	fs.Parse(args)

	if err := util.ValidateRequiredFlags(fs, []string{"id", "type", "reason", "date"}); err != nil {
		fmt.Printf("Validation error: %v\n", err)
		fs.Usage()
		os.Exit(1)
	}

	// แปลง string -> time.Time
	leaveDate, err := time.Parse("2006-01-02", *leaveDateStr)
	if err != nil {
		fmt.Printf("Invalid date format: %v\n", err)
		os.Exit(1)
	}

	// เปิด database และเตรียม facade
	db := util.OpenDatabase(*util.DatabasePath)
	hrFacade := controller.NewHRFacade(db)

	// สร้างคำขอลา
	request := hrModel.NewRequestLeaveBuilder().
		WithStudentID(*studentID).
		WithLeaveType(*leaveType).
		WithReason(*reason).
		WithLeaveDate(leaveDate).
		Build()

	
	if err := hrFacade.SubmitLeaveRequest(request); err != nil {
		fmt.Printf("Failed to submit leave request: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Leave request submitted successfully.")
}
