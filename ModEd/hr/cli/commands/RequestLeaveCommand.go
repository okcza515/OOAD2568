package commands

import (
	"ModEd/hr/controller"
	"ModEd/hr/util"
	"flag"
	"fmt"

	"gorm.io/gorm"
)

func requestLeave(target string, args []string, tx *gorm.DB) error {
	fs := flag.NewFlagSet("requestLeave", flag.ExitOnError)

	idUsage := "User ID"
	if target == "student" {
		idUsage = "Student ID"
	} else if target == "instructor" {
		idUsage = "Instructor ID"
	}

	id := fs.String("id", "", idUsage)
	leaveType := fs.String("type", "", "Type of leave (e.g. sick, personal)")
	reason := fs.String("reason", "", "Reason for leave")
	leaveDateStr := fs.String("date", "", "Leave date (YYYY-MM-DD)")
	fs.Parse(args)

	validator := util.NewValidationChain(fs)
	validator.Field("id").Required().Length(11).Regex(`^[0-9]{11}$`)
	validator.Field("type").Required()
	validator.Field("reason").Required()
	validator.Field("date").Required().Regex(`^\d{4}-\d{2}-\d{2}$`)
	validator.Validate()
	err := validator.Validate()
	if err != nil {
		fs.Usage()
		return fmt.Errorf("validation error: %v", err)
	}

	if target == "student" {
		err = controller.SubmitStudentLeaveRequest(tx, *id, *leaveType, *reason, *leaveDateStr)
	} else if target == "instructor" {
		err = controller.SubmitInstructorLeaveRequest(tx, *id, *leaveType, *reason, *leaveDateStr)
	} else {
		return fmt.Errorf("internal error: invalid target '%s' for requestLeave", target)
	}

	if err != nil {
		return err
	}

	fmt.Println("Leave request submitted successfully.")
	return nil
}
