package handler

import (
	"ModEd/core"
	"ModEd/core/validation"
	"ModEd/hr/controller"
	"fmt"
)

type AddInstructorStrategy struct {
	instructorController *controller.InstructorHRController
}

func NewAddInstructorStrategy(instructorCtrl *controller.InstructorHRController) *AddInstructorStrategy {
	return &AddInstructorStrategy{instructorController: instructorCtrl}
}

func (handler AddInstructorStrategy) Execute() error {
	validator := validation.NewValidationChain(core.GetUserInput)

	instructorCode := validator.Field(validation.FieldConfig{Name: "instructorCode", Prompt: "Enter instructor code: "}).Required().GetInput()
	firstName := validator.Field(validation.FieldConfig{Name: "firstName", Prompt: "Enter instructor first name: "}).Required().GetInput()
	lastName := validator.Field(validation.FieldConfig{Name: "lastName", Prompt: "Enter instructor last name: "}).Required().GetInput()
	email := validator.Field(validation.FieldConfig{Name: "email", Prompt: "Enter instructor email: "}).Required().IsEmail().GetInput()
	startDate := validator.Field(validation.FieldConfig{Name: "startDate", Prompt: "Enter instructor start date (YYYY-MM-DD): "}).Required().IsDate().GetInput()
	department := validator.Field(validation.FieldConfig{Name: "department", Prompt: "Enter instructor department: "}).Required().GetInput()
	gender := validator.Field(validation.FieldConfig{Name: "gender", Prompt: "Enter instructor gender (e.g., Male, Female, Other): "}).Required().GetInput()
	citizenID := validator.Field(validation.FieldConfig{Name: "citizenID", Prompt: "Enter instructor citizen ID (13 digits): "}).Required().IsAllDigits().Length(13).GetInput()
	phoneNumber := validator.Field(validation.FieldConfig{Name: "phoneNumber", Prompt: "Enter instructor phone number (10 digits, e.g., 0812345678): "}).Required().IsPhoneNumber().GetInput()
	salary := validator.Field(validation.FieldConfig{Name: "salary", Prompt: "Enter instructor salary: "}).Required().GetParsedNumber()

	academicPos := validator.Field(validation.FieldConfig{Name: "academicPos", Prompt: "Enter instructor academic position: "}).Required().GetInput()
	departmentPos := validator.Field(validation.FieldConfig{Name: "departmentPos", Prompt: "Enter instructor department position: "}).Required().GetInput()

	err := handler.instructorController.AddInstructor(
		instructorCode,
		firstName,
		lastName,
		email,
		startDate,
		department,
		gender,
		citizenID,
		phoneNumber,
		salary,
		academicPos,
		departmentPos,
	)

	if err != nil {
		return fmt.Errorf("failed to add instructor: %w", err)
	}

	fmt.Println("Instructor is added successfully!")

	return nil
}
