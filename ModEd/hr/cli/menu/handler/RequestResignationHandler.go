package handler

import (
	"ModEd/core"
	"ModEd/core/validation"
	"fmt"
)

type ControllerReqResignFunc func(
	id string,
	reason string,
) error

type RequestResignationHandler struct {
	controllerFunc ControllerReqResignFunc
}

func NewRequestResignationHandlerStrategy(controllerFunc ControllerReqResignFunc) *RequestResignationHandler {
	return &RequestResignationHandler{controllerFunc: controllerFunc}
}

func (handler RequestResignationHandler) Execute() error {
	validator := validation.NewValidationChain(core.GetUserInput)

	id := validator.Field(validation.FieldConfig{Name: "id", Prompt: "Enter ID: "}).Required().GetInput()
	reason := validator.Field(validation.FieldConfig{Name: "reason", Prompt: "Enter reason: "}).Required().GetInput()

	err := handler.controllerFunc(id, reason)
	if err != nil {
		return fmt.Errorf("failed to submit resignation request: %v", err)
	}

	fmt.Println("Resignation request submitted successfully!")
	return nil
}
