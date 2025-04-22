package commands

import (
	"ModEd/hr/controller"
	hrModel "ModEd/hr/model"
	"ModEd/hr/util"
	"flag"
	"fmt"

	"gorm.io/gorm"
)

// usage: go run hr/cli/HumanResourceCLI.go request student leave -id="66050001" -type="sick" -reason="ไม่สบาย" -date="2025-04-20"
func requestLeaveStudent(args []string, tx *gorm.DB) error {
	// fmt.Printf("asdjoasjdojaodjsaojdoasjodsjodasj")
	fs := flag.NewFlagSet("requestLeave", flag.ExitOnError)
	studentID := fs.String("id", "", "Student ID")
	leaveType := fs.String("type", "", "Type of leave (e.g. sick, personal)")
	reason := fs.String("reason", "", "Reason for leave")
	leaveDateStr := fs.String("date", "", "Leave date (YYYY-MM-DD)")
	fs.Parse(args)

	if err := util.ValidateRequiredFlags(fs, []string{"id", "type", "reason", "date"}); err != nil {
		fs.Usage()
		return fmt.Errorf("validation error: %v", err)
	}

	db := util.OpenDatabase(*util.DatabasePath)
	tm := &util.TransactionManager{DB: tx}

	err := tm.Execute(func(tx *gorm.DB) error {
		hrFacade := controller.NewHRFacade(db)
		factory := &hrModel.RequestLeaveFactory{}
		req, err := factory.Create("student", *studentID, *leaveType, *reason, *leaveDateStr)
		if err != nil {
			return fmt.Errorf("failed to build leave request: %v", err)
		}

		if err := hrFacade.SubmitLeaveStudentRequest(req.(*hrModel.RequestLeaveStudent)); err != nil {
			return fmt.Errorf("failed to submit leave request: %v", err)
		}
		return nil
	})
	if err != nil {
		return err
	}

	fmt.Println("Leave request submitted successfully.")
	return nil
}
