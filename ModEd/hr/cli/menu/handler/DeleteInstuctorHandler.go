package handler

import (
	"ModEd/core"
	"ModEd/core/validation"
	"ModEd/hr/controller"
	"fmt"
)

type DeleteInstructorStrategy struct {
	instructorController *controller.InstructorHRController
}

func NewDeleteInstructorStrategy(instructorCtrl *controller.InstructorHRController) *DeleteInstructorStrategy {
	return &DeleteInstructorStrategy{instructorController: instructorCtrl}
}

func (handler DeleteInstructorStrategy) Execute() error {
	validator := validation.NewValidationChain(core.GetUserInput)

	instructorCode := validator.Field(validation.FieldConfig{
		Name:   "instructorCode",
		Prompt: "Enter instructor code : ",
	}).Required().GetInput()

	err := handler.instructorController.DeleteInstructor(instructorCode)
	if err != nil {
		return fmt.Errorf("failed to delete instructor: %w", err)
	}

	fmt.Println("Instructor deleted successfully!")
	return nil
}
