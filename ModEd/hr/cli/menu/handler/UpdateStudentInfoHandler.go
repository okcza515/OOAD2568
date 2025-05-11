package handler

import (
	"ModEd/core"
	"ModEd/core/validation"
	"ModEd/hr/controller"
	"fmt"
)

type UpdateStudentInfoStrategy struct {
	studentController *controller.StudentHRController
}

func NewUpdateStudentInfoStrategy(studentCtrl *controller.StudentHRController) *UpdateStudentInfoStrategy {
	return &UpdateStudentInfoStrategy{studentController: studentCtrl}
}

func (handler UpdateStudentInfoStrategy) Execute() error {
	validator := validation.NewValidationChain(core.GetUserInput)

	studentCode := validator.Field(validation.FieldConfig{Name: "studentCode", Prompt: "Enter student ID (11 digits): "}).Required().IsStudentCode().GetInput()
	firstName := validator.Field(validation.FieldConfig{Name: "firstName", Prompt: "Enter student first name: "}).Required().GetInput()
	lastName := validator.Field(validation.FieldConfig{Name: "lastName", Prompt: "Enter student last name: "}).Required().GetInput()
	email := validator.Field(validation.FieldConfig{Name: "email", Prompt: "Enter student email: "}).Required().IsEmail().GetInput()
	gender := validator.Field(validation.FieldConfig{Name: "gender", Prompt: "Enter student gender (e.g., Male, Female, Other): "}).Required().GetInput() // Consider .AllowedValues()
	citizenID := validator.Field(validation.FieldConfig{Name: "citizenID", Prompt: "Enter student citizen ID (13 digits): "}).Required().IsAllDigits().Length(13).GetInput()
	phoneNumber := validator.Field(validation.FieldConfig{Name: "phoneNumber", Prompt: "Enter student phone number (10 digits, e.g., 0812345678): "}).Required().IsPhoneNumber().GetInput()

	err := handler.studentController.UpdateStudentInfo(
		studentCode,
		firstName,
		lastName,
		email,
		gender,
		citizenID,
		phoneNumber,
	)
	
	if err != nil {
		return fmt.Errorf("failed to update student info: %w", err)
	}
	
	fmt.Println("Student updated successfully!")
	return nil
}
