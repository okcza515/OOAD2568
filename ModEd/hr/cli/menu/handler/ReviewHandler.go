package handler

import (
	"ModEd/core"
	"ModEd/core/validation"
	"fmt"
	"strings"

)

type ControllerReviewFunc func(requestID string, action string, reason string) error

type ReviewHandler struct {
	ControllerFunc ControllerReviewFunc
}

func NewReviewHandlerStrategy(controllerFunc ControllerReviewFunc) *ReviewHandler {
	return &ReviewHandler{ ControllerFunc: controllerFunc,}
}

func (handler ReviewHandler) Execute() error {
	validator := validation.NewValidationChain(core.GetUserInput)

	requestID := validator.Field(validation.FieldConfig{Name: "id", Prompt: "Enter Request ID:"}).Required().GetInput()
	action := validator.Field(validation.FieldConfig{Name: "action", Prompt: "Enter action (approve/reject):"}).Required().AllowedValues([]string{"approve", "reject"}).GetInput()

	var reason string
	if strings.ToLower(action) == "reject" {
		reason = validator.Field(validation.FieldConfig{Name: "reason", Prompt: "Enter rejection reason:"}).Required().GetInput()
	}

	err := handler.ControllerFunc(requestID, strings.ToLower(action), reason)
	if err != nil {
		return fmt.Errorf("failed to review request: %w", err)
	}

	fmt.Printf("Request '%s' was %sed successfully!\n", requestID, action)
	return nil
}
