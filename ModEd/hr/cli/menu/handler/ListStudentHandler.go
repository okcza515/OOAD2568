package handler

import (
	"ModEd/hr/controller"
	"ModEd/hr/util"
	"fmt"
)

type ListStudentStrategy struct {
	studentController *controller.StudentHRController
}

func NewListStudentStrategy(studentCtrl *controller.StudentHRController) *ListStudentStrategy {
	return &ListStudentStrategy{studentController: studentCtrl}
}

func (handler ListStudentStrategy) Execute() error {
	studentInfos, err := handler.studentController.GetAllStudents()
	if err != nil {
		return fmt.Errorf("error listing students: %v", err)
	}

	fmt.Println("Human Resource Student Info:")
	for _, s := range studentInfos {
		fmt.Printf("SID: %s | Name: %s %s | Gender: %s | CitizenID: %s | Phone: %s | Status: %s | Email: %s\n",
			s.StudentCode, s.FirstName, s.LastName, s.Gender, s.CitizenID, s.PhoneNumber, util.StatusToString(*s.Status), s.Email)
	}
	return nil
}
