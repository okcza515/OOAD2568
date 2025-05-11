package handler

import (
	"ModEd/core"
	"ModEd/core/validation"
	"ModEd/hr/controller"
	"fmt"
)

type DeleteStudentStrategy struct {
	studentController *controller.StudentHRController
}

func NewDeleteStudentStrategy(studentCtrl *controller.StudentHRController) *DeleteStudentStrategy {
	return &DeleteStudentStrategy{studentController: studentCtrl}
}

func (handler DeleteStudentStrategy) Execute() error {
	validator := validation.NewValidationChain(core.GetUserInput)

	studentCode := validator.Field(validation.FieldConfig{Name: "studentCode", Prompt: "Enter student ID (11 digits): "}).Required().IsStudentCode().GetInput()

	err := handler.studentController.DeleteStudent(studentCode)

	if err != nil {
		return fmt.Errorf("failed to delete student: %w", err)
	}

	fmt.Println("Student deleted successfully!")
	return nil
}
