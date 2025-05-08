package handler

import (
	"ModEd/core"
	"ModEd/core/validation"
	"ModEd/hr/controller"
	"fmt"
)

type AddStudentStrategy struct {
	studentController *controller.StudentHRController
}

func NewAddStudentStrategy(studentCtrl *controller.StudentHRController) *AddStudentStrategy {
	return &AddStudentStrategy{studentController: studentCtrl}
}

func (handler AddStudentStrategy) Execute() error {
	validator := validation.NewValidationChain(core.GetUserInput)

	studentCode := validator.Field(validation.FieldConfig{Name: "studentCode", Prompt: "Enter student ID (11 digits): "}).Required().IsStudentCode().GetInput()
	firstName := validator.Field(validation.FieldConfig{Name: "firstName", Prompt: "Enter student first name: "}).Required().GetInput()
	lastName := validator.Field(validation.FieldConfig{Name: "lastName", Prompt: "Enter student last name: "}).Required().GetInput()
	email := validator.Field(validation.FieldConfig{Name: "email", Prompt: "Enter student email: "}).Required().IsEmail().GetInput()
	startDate := validator.Field(validation.FieldConfig{Name: "startDate", Prompt: "Enter student start date (YYYY-MM-DD HH:MM:SS): "}).Required().IsDateTime().GetInput()
	birthDate := validator.Field(validation.FieldConfig{Name: "birthDate", Prompt: "Enter student birth date (YYYY-MM-DD HH:MM:SS): "}).Required().IsDateTime().GetInput()
	program := validator.Field(validation.FieldConfig{Name: "program", Prompt: "Enter student program: "}).Required().GetInput()
	department := validator.Field(validation.FieldConfig{Name: "department", Prompt: "Enter student department: "}).Required().GetInput()
	status := validator.Field(validation.FieldConfig{Name: "status", Prompt: "Enter student status (e.g., Active, Graduated): "}).Required().GetInput()   // Consider .AllowedValues()
	gender := validator.Field(validation.FieldConfig{Name: "gender", Prompt: "Enter student gender (e.g., Male, Female, Other): "}).Required().GetInput() // Consider .AllowedValues()
	citizenID := validator.Field(validation.FieldConfig{Name: "citizenID", Prompt: "Enter student citizen ID (13 digits): "}).Required().IsAllDigits().Length(13).GetInput()
	phoneNumber := validator.Field(validation.FieldConfig{Name: "phoneNumber", Prompt: "Enter student phone number (10 digits, e.g., 0812345678): "}).Required().IsPhoneNumber().GetInput()
	advisorCode := validator.Field(validation.FieldConfig{Name: "advisorCode", Prompt: "Enter student advisor code: "}).Required().GetInput()

	err := handler.studentController.AddStudent(
		studentCode,
		firstName,
		lastName,
		email,
		startDate,
		birthDate,
		program,
		department,
		status,
		gender,
		citizenID,
		phoneNumber,
		advisorCode,
	)

	if err != nil {
		return fmt.Errorf("failed to add student: %w", err)
	}

	fmt.Println("Student is added successfully!")

	return nil
}
