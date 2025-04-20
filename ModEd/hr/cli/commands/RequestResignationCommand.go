package commands

import (
	"ModEd/hr/controller"
	"ModEd/hr/model"
	"ModEd/hr/util"
	"flag"
	"fmt"
	"strings"

	"gorm.io/gorm"
)

func (c *RequestResignationCommand) Execute(args []string, tx *gorm.DB) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: requestResignation {student|instructor} [options]")
	}

	target := strings.ToLower(args[0])
	switch target {
	case "student":
		return RequestResignationStudent(args[1:], tx)
	case "instructor":
		return RequestResignationInstructor(args[1:], tx)
	default:
		return fmt.Errorf("unknown requestResignation target: %s", target)
	}
}

func RequestResignationStudent(args []string, tx *gorm.DB) error {
	fs := flag.NewFlagSet("requestResignation", flag.ExitOnError)
	ID := fs.String("id", "", "ID")
	reason := fs.String("reason", "", "Reason for resignation")

	if err := fs.Parse(args); err != nil {
		return fmt.Errorf("failed to parse flags: %v", err)
	}

	if err := util.ValidateRequiredFlags(fs, []string{"id", "reason"}); err != nil {
		fs.Usage()
		return fmt.Errorf("validation error: %v", err)
	}

	tm := &util.TransactionManager{DB: tx}
	return tm.Execute(func(tx *gorm.DB) error {
		hrFacade := controller.NewHRFacade(tx)

		request := model.NewRequestResignationStudentBuilder().
			WithStudentID(*ID).
			WithReason(*reason).
			Build()

		if err := hrFacade.SubmitResignationStudentRequest(request); err != nil {
			return fmt.Errorf("failed to submit resignation request: %v", err)
		}

		fmt.Println("Resignation request submitted successfully.")
		return nil
	})
}

func RequestResignationInstructor(args []string, tx *gorm.DB) error {
	fs := flag.NewFlagSet("requestResignation", flag.ExitOnError)
	ID := fs.String("id", "", "ID")
	reason := fs.String("reason", "", "Reason for resignation")

	if err := fs.Parse(args); err != nil {
		return fmt.Errorf("failed to parse flags: %v", err)
	}

	if err := util.ValidateRequiredFlags(fs, []string{"reason"}); err != nil {
		fs.Usage()
		return fmt.Errorf("validation error: %v", err)
	}

	tm := &util.TransactionManager{DB: tx}
	return tm.Execute(func(tx *gorm.DB) error {
		hrFacade := controller.NewHRFacade(tx)

		request := model.NewRequestResignationInstructorBuilder().
			WithInstructorID(*ID).
			WithReason(*reason).
			Build()

		if err := hrFacade.SubmitResignationInstructorRequest(request); err != nil {
			return fmt.Errorf("failed to submit resignation request: %v", err)
		}

		fmt.Println("Resignation request submitted successfully.")
		return nil
	})
}
