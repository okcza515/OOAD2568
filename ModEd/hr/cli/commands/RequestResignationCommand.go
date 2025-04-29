package commands

import (
	"ModEd/hr/controller"
	"ModEd/hr/util"
	"flag"
	"fmt"

	"gorm.io/gorm"
)

// usage: go run hr/cli/HumanResourceCLI.go request student resign -id="66050001" -reason=""
// usage (instructor): go run hr/cli/HumanResourceCLI.go request instructor resign -id="I0001" -reason=""
func requestResignation(target string, args []string, tx *gorm.DB) error {
	fs := flag.NewFlagSet("requestResignation", flag.ExitOnError)

	idUsage := "User ID"
	if target == "student" {
		idUsage = "Student ID"
	} else if target == "instructor" {
		idUsage = "Instructor ID"
	}

	id := fs.String("id", "", idUsage)
	reason := fs.String("reason", "", "Reason for resignation")

	if err := fs.Parse(args); err != nil {
		return fmt.Errorf("failed to parse flags: %v", err)
	}

	validationChain := util.NewValidationChain(fs).
		Required("id").
		Required("reason")

	if target == "student" {
		validationChain.Length("id", 11)
	} else if target == "instructor" {
	}

	if err := validationChain.Validate(); err != nil {
		fs.Usage()
		return fmt.Errorf("validation error: %v", err)
	}

	var err error

	if target == "student" {
		err = controller.SubmitResignationStudent(tx, *id, *reason)
		if err != nil {
			return fmt.Errorf("failed to submit student resignation request: %v", err)
		}
		fmt.Println("Student resignation request submitted successfully.")
	} else if target == "instructor" {
		err = controller.SubmitResignationInstructor(tx, *id, *reason)
		if err != nil {
			return fmt.Errorf("failed to submit instructor resignation request: %v", err)
		}
		fmt.Println("Instructor resignation request submitted successfully.")
	} else {
		return fmt.Errorf("internal error: invalid target '%s' for requestResignation", target)
	}

	return nil
}
