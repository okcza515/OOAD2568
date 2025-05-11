package handler

import (
	"ModEd/core"
	"ModEd/core/validation"
	"ModEd/hr/controller"
	"fmt"
)

func handleLeaveRequest(target string, studentLeaveCtrl *controller.LeaveStudentHRController, instructorLeaveCtrl *controller.LeaveInstructorHRController) error {
	validator := validation.NewValidationChain(core.GetUserInput)

	id := validator.Field(validation.FieldConfig{Name: "id", Prompt: "Enter ID: "}).Required().GetInput()
	leaveType := validator.Field(validation.FieldConfig{Name: "leaveType", Prompt: "Enter leave type: "}).Required().GetInput()
	reason := validator.Field(validation.FieldConfig{Name: "reason", Prompt: "Enter reason: "}).Required().GetInput()
	leaveDateStr := validator.Field(validation.FieldConfig{Name: "leaveDate", Prompt: "Enter leave date (YYYY-MM-DD): "}).Required().GetInput()

	operations := map[string]func() error{
		"student": func() error {
			err := studentLeaveCtrl.SubmitStudentLeaveRequest(id, leaveType, reason, leaveDateStr)
			if err != nil {
				return err
			}
			return nil
		},
		"instructor": func() error {
			err := instructorLeaveCtrl.SubmitInstructorLeaveRequest(id, leaveType, reason, leaveDateStr)
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

func handleResignationRequest(target string, studentResignCtrl *controller.ResignationStudentHRController, instructorResignCtrl *controller.ResignationInstructorHRController) error {
	validator := validation.NewValidationChain(core.GetUserInput)

	id := validator.Field(validation.FieldConfig{Name: "id", Prompt: "Enter ID: "}).Required().GetInput()
	reason := validator.Field(validation.FieldConfig{Name: "reason", Prompt: "Enter reason: "}).Required().GetInput()

	operations := map[string]func() error{
		"student": func() error {
			err := studentResignCtrl.SubmitResignationStudent(id, reason)
			if err != nil {
				return err
			}
			return nil
		},
		"instructor": func() error {
			err := instructorResignCtrl.SubmitResignationInstructor(id, reason)
			if err != nil {
				return err
			}
			return nil
		},
	}

	operation, exists := operations[target]
	if !exists {
		return fmt.Errorf("internal error: invalid target '%s' for handleResignationRequest", target)
	}

	if err := operation(); err != nil {
		return fmt.Errorf("failed to submit %s resignation request: %v", target, err)
	}

	fmt.Printf("%s resignation request submitted successfully!\n", target)
	return nil
}
