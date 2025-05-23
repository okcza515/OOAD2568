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

	validator := util.NewValidationChain(fs)
	validator.Field("type").Required()
	validator.Field("reason").Required()
	validator.Field("date").Required().IsDate()

	operations := map[string]func() error{
		"student": func() error {
			validator.Field("id").Required().IsStudentID()
			if err := validator.Validate(); err != nil {
				fs.Usage()
				return fmt.Errorf("validation error for student leave: %v", err)
			}
			leaveStudentController := controller.NewLeaveStudentHRController(tx)
			return leaveStudentController.SubmitStudentLeaveRequest(*id, *leaveType, *reason, *leaveDateStr)
		},
		"instructor": func() error {
			validator.Field("id").Required().IsInstructorID()
			if err := validator.Validate(); err != nil {
				fs.Usage()
				return fmt.Errorf("validation error for instructor leave: %v", err)
			}
			leaveInstructorController := controller.NewLeaveInstructorHRController(tx)
			return leaveInstructorController.SubmitInstructorLeaveRequest(*id, *leaveType, *reason, *leaveDateStr)
		},
	}

	operation, exists := operations[target]
	if !exists {
		return fmt.Errorf("internal error: invalid target '%s' for handleLeaveRequest", target)
	}

	if err := operation(); err != nil {
		return fmt.Errorf("failed to submit %s leave request: %v", target, err)
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

	validator := util.NewValidationChain(fs)

	operations := map[string]func() error{
		"student": func() error {
			validator.Field("id").Required().IsStudentID()
			if err := validator.Validate(); err != nil {
				fs.Usage()
				return fmt.Errorf("validation error for student: %v", err)
			}
			controller := controller.NewResignationStudentHRController(tx)
			return controller.SubmitResignationStudent(*id, *reason)
		},
		"instructor": func() error {
			validator.Field("id").Required().IsInstructorID()
			if err := validator.Validate(); err != nil {
				fs.Usage()
				return fmt.Errorf("validation error for instructor: %v", err)
			}
			controller := controller.NewResignationInstructorHRController(tx)
			return controller.SubmitResignationInstructor(*id, *reason)
		},
	}

	operation, exists := operations[target]
	if !exists {
		return fmt.Errorf("internal error: invalid target '%s' for handleResignationRequest", target)
	}

	if err := operation(); err != nil {
		return fmt.Errorf("failed to submit resignation request: %v", err)
	}

	fmt.Printf("%s resignation request submitted successfully.\n", target)
	return nil
}

// usage: go run hr/cli/HumanResourceCLI.go request instructor raise -id="66050001" -amount=10000 -reason="ดีมาก"
// usage (student): go run hr/cli/HumanResourceCLI.go request student raise -id="66050001" -amount=10000 -reason="ดีมาก"
func handleRaiseRequest(target string, args []string, tx *gorm.DB) error {
	commandName := fmt.Sprintf("request-%s-raise", target)
	fs := flag.NewFlagSet(commandName, flag.ExitOnError)

	idUsage := fmt.Sprintf("%s ID", target)
	id := fs.String("id", "", idUsage)
	amount := fs.Float64("amount", 0, "Raise amount")
	reason := fs.String("reason", "", "Reason for raise")

	if err := fs.Parse(args); err != nil {
		return fmt.Errorf("failed to parse flags: %v", err)
	}

	validator := util.NewValidationChain(fs)
	validator.Field("amount").Required()
	validator.Field("reason").Required()

	operations := map[string]func() error{
		"student": func() error {
			validator.Field("id").Required().IsStudentID()
			if err := validator.Validate(); err != nil {
				fs.Usage()
				return fmt.Errorf("validation error for student raise: %v", err)
			}
			controller := controller.CreateRaiseInstructorHRController(tx)
			return controller.SubmitRaiseRequest(*id, *amount, *reason)
		},
		"instructor": func() error {
			validator.Field("id").Required().IsInstructorID()
			if err := validator.Validate(); err != nil {
				fs.Usage()
				return fmt.Errorf("validation error for instructor raise: %v", err)
			}
			controller := controller.NewRaiseHRController(tx)
			return controller.SubmitRaiseRequest(*id, *amount, *reason)
		},
	}

	operation, exists := operations[target]
	if !exists {
		return fmt.Errorf("internal error: invalid target '%s' for handleRaiseRequest", target)
	}

	if err := operation(); err != nil {
		return fmt.Errorf("failed to submit raise request: %v", err)
	}

	fmt.Printf("%s raise request submitted successfully.\n", target)
	return nil
}
