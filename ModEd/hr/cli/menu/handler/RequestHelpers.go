package handler

import (
	"ModEd/core"
	"ModEd/core/validation"
	"ModEd/hr/controller"
	"fmt"
)

func handleLeaveRequest(target string, studentCtrl *controller.LeaveStudentHRController, instructorCtrl *controller.LeaveInstructorHRController) error {
	validator := validation.NewValidationChain(core.GetUserInput)

	id := validator.Field(validation.FieldConfig{Name: "id", Prompt: "Enter ID: "}).Required().GetInput()
	leaveType := validator.Field(validation.FieldConfig{Name: "leaveType", Prompt: "Enter leave type: "}).Required().GetInput()
	reason := validator.Field(validation.FieldConfig{Name: "reason", Prompt: "Enter reason: "}).Required().GetInput()
	leaveDateStr := validator.Field(validation.FieldConfig{Name: "leaveDate", Prompt: "Enter leave date (YYYY-MM-DD): "}).Required().GetInput()

	operations := map[string]func() error{
		"student": func() error {
			err := studentCtrl.SubmitStudentLeaveRequest(id, leaveType, reason, leaveDateStr)
			if err != nil {
				return err
			}
			return nil
		},
		"instructor": func() error {
			err := instructorCtrl.SubmitInstructorLeaveRequest(id, leaveType, reason, leaveDateStr)
			if err != nil {
				return err
			}
			return nil
		},
	}

	operation, exists := operations[target]
	if !exists {
		return fmt.Errorf("internal error: invalid target '%s' for handleLeaveRequest", target)
	}

	if err := operation(); err != nil {
		return fmt.Errorf("failed to submit %s leave request: %v", target, err)
	}

	fmt.Printf("%s leave request submitted successfully!\n", target)
	return nil
}
