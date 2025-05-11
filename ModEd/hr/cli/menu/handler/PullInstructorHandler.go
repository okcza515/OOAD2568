package handler

import (
	"ModEd/hr/controller"
	"fmt"
)

type PullInstructorHandler struct {
	instructorController *controller.InstructorHRController
}

func NewPullInstructorHandlerStrategy(instructorCtrl *controller.InstructorHRController) *PullInstructorHandler {
	return &PullInstructorHandler{instructorController: instructorCtrl}
}

func (handler PullInstructorHandler) Execute() error {
	if err := handler.instructorController.MigrateInstructorRecords(); err != nil {
		return fmt.Errorf("failed to pull instructor record into instructorinfo: %w", err)
	}

	fmt.Println("Instructor record pulled successfully into InstructorInfo table.")
	return nil
}
