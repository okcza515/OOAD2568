package handler

import (
	"ModEd/core"
	"ModEd/core/validation"
	"ModEd/hr/controller"
	"fmt"
)

type UpdateInstructorInfoStrategy struct {
	instructorController *controller.InstructorHRController
}

func NewUpdateInstructorInfoStrategy(instructorCtrl *controller.InstructorHRController) *UpdateInstructorInfoStrategy {
	return &UpdateInstructorInfoStrategy{instructorController: instructorCtrl}
}

func (handler UpdateInstructorInfoStrategy) Execute() error {
	validator := validation.NewValidationChain(core.GetUserInput)

	instructorID := validator.Field(validation.FieldConfig{Name: "instructorID", Prompt: "Enter instructor ID (11 digits): "}).Required().IsAllDigits().GetInput()
	firstName := validator.Field(validation.FieldConfig{Name: "firstName", Prompt: "Enter instructor first name: "}).Required().GetInput()
	lastName := validator.Field(validation.FieldConfig{Name: "lastName", Prompt: "Enter instructor last name: "}).Required().GetInput()
	email := validator.Field(validation.FieldConfig{Name: "email", Prompt: "Enter instructor email: "}).Required().IsEmail().GetInput()
	gender := validator.Field(validation.FieldConfig{Name: "gender", Prompt: "Enter instructor gender (e.g., Male, Female, Other): "}).Required().GetInput()
	citizenID := validator.Field(validation.FieldConfig{Name: "citizenID", Prompt: "Enter instructor citizen ID (13 digits): "}).Required().IsAllDigits().Length(13).GetInput()
	phoneNumber := validator.Field(validation.FieldConfig{Name: "phoneNumber", Prompt: "Enter instructor phone number (10 digits, e.g., 0812345678): "}).Required().IsPhoneNumber().GetInput()
	academicPos := validator.Field(validation.FieldConfig{Name: "academicPos", Prompt: "Enter academic position: "}).Required().GetInput()
	departmentPos := validator.Field(validation.FieldConfig{Name: "departmentPos", Prompt: "Enter department position: "}).Required().GetInput()

	err := handler.instructorController.UpdateInstructorInfo(
		instructorID,
		firstName,
		lastName,
		email,
		gender,
		citizenID,
		phoneNumber,
		academicPos,
		departmentPos,
	)
	if err != nil {
		return fmt.Errorf("failed to update instructor info: %w", err)
	}

	fmt.Println("Instructor updated successfully!")
	return nil
}
