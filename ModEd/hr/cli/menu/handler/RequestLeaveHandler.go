package handler

import (
	"ModEd/core"
	"ModEd/core/validation"
	"fmt"
)

type ControllerReqLeaveFunc func(
	id string,
	leaveType string,
	reason string,
	leaveDateStr string,
) error

type RequestLeaveHandler struct {
	controllerFunc ControllerReqLeaveFunc
}

func NewRequestLeaveHandlerStrategy(controllerFunc ControllerReqLeaveFunc) *RequestLeaveHandler {
	return &RequestLeaveHandler{controllerFunc: controllerFunc}
}

func (handler RequestLeaveHandler) Execute() error {
	validator := validation.NewValidationChain(core.GetUserInput)

	id := validator.Field(validation.FieldConfig{Name: "id", Prompt: "Enter ID: "}).Required().GetInput()
	leaveType := validator.Field(validation.FieldConfig{Name: "leaveType", Prompt: "Enter leave type: "}).Required().GetInput()
	reason := validator.Field(validation.FieldConfig{Name: "reason", Prompt: "Enter reason: "}).Required().GetInput()
	leaveDateStr := validator.Field(validation.FieldConfig{Name: "leaveDate", Prompt: "Enter leave date (YYYY-MM-DD): "}).Required().IsDate().GetInput()

	err := handler.controllerFunc(id, leaveType, reason, leaveDateStr)
	if err != nil {
		return fmt.Errorf("failed to submit leave request: %v", err)
	}

	fmt.Println("Leave request submitted successfully!")
	return nil
}
