package handler

import (
	"ModEd/core"
	"ModEd/core/validation"
	"fmt"
	"strings"

	"gorm.io/gorm"
)

type ControllerReviewFunc func(tx *gorm.DB, requestID string, action string, reason string) error

func HandleReviewRequest(
	tx *gorm.DB,
	entityType string, 
	controllerFunc ControllerReviewFunc,
) error {
	validator := validation.NewValidationChain(core.GetUserInput)

	requestID := validator.Field(validation.FieldConfig{Name: "id", Prompt: "Enter Request ID:"}).Required().GetInput()
	action := validator.Field(validation.FieldConfig{Name: "action", Prompt: "Enter action (approve/reject):"}).Required().AllowedValues([]string{"approve", "reject"}).GetInput()

	var reason string
	if strings.ToLower(action) == "reject" {
		reason = validator.Field(validation.FieldConfig{Name: "reason", Prompt: "Enter rejection reason:"}).Required().GetInput()
	}
	err := controllerFunc(tx, requestID, strings.ToLower(action), reason)
	if err != nil {
		return fmt.Errorf("failed to review %s request: %v", entityType, err)
	}

	fmt.Printf("%s request '%s' %sed successfully!\n", entityType, requestID, action)
	return nil
}
