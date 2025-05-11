package handler

import (
	"ModEd/core"
	"ModEd/core/validation"
	"ModEd/hr/controller"
	"fmt"
)

type RequestRaiseStrategy struct {
	requestRaiseController *controller.RaiseHRController
}

func NewRequestRaiseHandlerStrategy(requestRaiseCtrl *controller.RaiseHRController) *RequestRaiseStrategy {
	return &RequestRaiseStrategy{requestRaiseController: requestRaiseCtrl}
}

func (handler RequestRaiseStrategy) Execute() error {
	validator := validation.NewValidationChain(core.GetUserInput)

	id := validator.Field(validation.FieldConfig{Name: "id", Prompt: "Enter your ID: "}).Required().GetInput()
	reason := validator.Field(validation.FieldConfig{Name: "reason", Prompt: "Enter reason for raise request: "}).Required().GetInput()
	amount := validator.Field(validation.FieldConfig{Name: "amount", Prompt: "Enter requested raise amount: "}).Required().GetParsedNumber()

	err := handler.requestRaiseController.SubmitRaiseRequest(id, amount, reason)
	if err != nil {
		return fmt.Errorf("failed to submit raise request: %v", err)
	}

	fmt.Println("Raise request submitted successfully!")
	return nil
}
