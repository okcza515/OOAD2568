package handler

import (
	"ModEd/hr/controller"
	"fmt"
)

type PullStudentHandler struct {
	studentController *controller.StudentHRController
}

func NewPullStudentHandlerStrategy(studentCtrl *controller.StudentHRController) *PullStudentHandler {
	return &PullStudentHandler{studentController: studentCtrl}
}

func (handler PullStudentHandler) Execute() error {
	if err := handler.studentController.MigrateStudentRecords(); err != nil {
		return fmt.Errorf("failed to pull student record into studentinfo: %w", err)
	}

	fmt.Println("Student record pulled successfully into StudentInfo table.")
	return nil
}
