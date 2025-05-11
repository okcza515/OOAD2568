package handler

import (
	"ModEd/hr/controller"
	"fmt"

	"gorm.io/gorm"
)

type PullStudentHandler struct {
	tx *gorm.DB
}

func NewPullStudentHandlerStrategy(tx *gorm.DB) *PullStudentHandler {
	return &PullStudentHandler{tx: tx}
}

func (handler PullStudentHandler) Execute() error {
	studentController := controller.NewStudentHRController(handler.tx)
	if err := studentController.MigrateStudentRecords(); err != nil {
		return fmt.Errorf("failed to pull student record into studentinfo: %w", err)
	}

	fmt.Println("Student record pulled successfully into StudentInfo table.")
	return nil
}
