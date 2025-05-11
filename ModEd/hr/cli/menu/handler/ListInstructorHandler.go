package handler

import (
	"ModEd/hr/controller"
	"fmt"
)

type ListInstructorHandler struct {
	instructorController *controller.InstructorHRController
}

func NewListInstructorHandler(instructorCtrl *controller.InstructorHRController) *ListInstructorHandler {
	return &ListInstructorHandler{instructorController: instructorCtrl}
}

func (handler ListInstructorHandler) Execute() error {
	instructorInfos, err := handler.instructorController.GetAllInstructors()
	if err != nil {
		return fmt.Errorf("error listing instructors: %w", err)
	}

	fmt.Println("Human Resource Instructor Info:")
	for _, i := range instructorInfos {
		fmt.Printf("IID: %s | Name: %s %s | Gender: %s | Email: %s | Department: %s | Phone: %s\n", i.InstructorCode, i.FirstName, i.LastName, i.Gender, i.Email, *i.Department, i.PhoneNumber)
	}

	return nil
}
