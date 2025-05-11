package handler

import (
	"ModEd/hr/controller"
	"fmt"

	"gorm.io/gorm"
)

type PullInstructorHandler struct {
	tx *gorm.DB
}

func NewPullInstructorHandlerStrategy(tx *gorm.DB) *PullInstructorHandler {
	return &PullInstructorHandler{tx: tx}
}

func (handler PullInstructorHandler) Execute() error {
	instructorController := controller.NewInstructorHRController(handler.tx)
	if err := instructorController.MigrateInstructorRecords(); err != nil {
		return fmt.Errorf("failed to pull instructor record into instructorinfo: %w", err)
	}

	fmt.Println("Instructor record pulled successfully into InstructorInfo table.")
	return nil
}
