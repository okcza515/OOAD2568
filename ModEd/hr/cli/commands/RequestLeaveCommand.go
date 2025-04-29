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

	// Validation remains the same
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

	if target == "student" {
		err = hrFacade.SubmitStudentLeaveRequest(tx, *id, *leaveType, *reason, *leaveDateStr)
	} else if target == "instructor" {
		err = hrFacade.SubmitInstructorLeaveRequest(tx, *id, *leaveType, *reason, *leaveDateStr)
	} else {
		return fmt.Errorf("internal error: invalid target '%s' for requestLeave", target)
	}

	if err != nil {
		return err
	}

	fmt.Println("Leave request submitted successfully.")
	return nil
}
