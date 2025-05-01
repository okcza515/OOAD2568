package commands

import (
	"ModEd/hr/controller"
	"ModEd/hr/util"
	"flag"
	"fmt"

	"gorm.io/gorm"
)

func handleLeaveRequest(target string, args []string, tx *gorm.DB) error {
	commandName := fmt.Sprintf("request-%s-leave", target)
	fs := flag.NewFlagSet(commandName, flag.ExitOnError)

	idUsage := fmt.Sprintf("%s ID", target)
	id := fs.String("id", "", idUsage)
	leaveType := fs.String("type", "", "Type of leave (e.g. sick, personal)")
	reason := fs.String("reason", "", "Reason for leave")
	leaveDateStr := fs.String("date", "", "Leave date (YYYY-MM-DD)")

	if err := fs.Parse(args); err != nil {
		return fmt.Errorf("failed to parse flags: %v", err)
	}

	var submitErr error
	validator := util.NewValidationChain(fs)

	validator.Field("type").Required()
	validator.Field("reason").Required()
	validator.Field("date").Required().IsDate()

	switch target {
	case "student":
		validator.Field("id").Required().IsStudentID()
		err := validator.Validate()
		if err != nil {
			fs.Usage()
			return fmt.Errorf("validation error for student leave: %v", err)
		}
		submitErr = controller.SubmitStudentLeaveRequest(tx, *id, *leaveType, *reason, *leaveDateStr)

	case "instructor":
		// Instructor Code in demo data is different from real instructor ID in LEB2.
		validator.Field("id").Required().IsInstructorID()

		err := validator.Validate()
		if err != nil {
			fs.Usage()
			return fmt.Errorf("validation error for instructor leave: %v", err)
		}
		submitErr = controller.SubmitInstructorLeaveRequest(tx, *id, *leaveType, *reason, *leaveDateStr)

	default:
		return fmt.Errorf("internal error: invalid target '%s' for handleLeaveRequest", target)
	}

	if submitErr != nil {
		return fmt.Errorf("failed to submit %s leave request: %v", target, submitErr)
	}

	fmt.Printf("%s leave request submitted successfully.\n", target)
	return nil
}

// usage: go run hr/cli/HumanResourceCLI.go request student resign -id="66050001" -reason=""
// usage (instructor): go run hr/cli/HumanResourceCLI.go request instructor resign -id="I0001" -reason=""
func handleResignationRequest(target string, args []string, tx *gorm.DB) error {
	commandName := fmt.Sprintf("request-%s-resign", target)
	fs := flag.NewFlagSet(commandName, flag.ExitOnError)

	idUsage := fmt.Sprintf("%s ID", target)
	id := fs.String("id", "", idUsage)
	reason := fs.String("reason", "", "Reason for resignation (optional)")

	if err := fs.Parse(args); err != nil {
		return fmt.Errorf("failed to parse flags: %v", err)
	}

	var submitErr error
	validator := util.NewValidationChain(fs)

	switch target {
	case "student":
		validator.Field("id").Required().IsStudentID()
		err := validator.Validate()
		if err != nil {
			fs.Usage()
			return fmt.Errorf("validation error for student: %v", err)
		}
		submitErr = controller.SubmitResignationStudent(tx, *id, *reason)

	case "instructor":
		// Instructor Code in demo data is different from real instructor ID in LEB2.
		validator.Field("id").Required().IsInstructorID()
		err := validator.Validate()
		if err != nil {
			fs.Usage()
			return fmt.Errorf("validation error for instructor: %v", err)
		}
		submitErr = controller.SubmitResignationInstructor(tx, *id, *reason)

	default:
		return fmt.Errorf("internal error: invalid target '%s' for handleResignationRequest", target)
	}

	if submitErr != nil {
		return fmt.Errorf("failed to submit resignation request: %v", submitErr)
	}

	fmt.Printf("%s resignation request submitted successfully.\n", target)
	return nil
}
