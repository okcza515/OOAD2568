package handler

import (
	"ModEd/hr/controller"

	"gorm.io/gorm"
)

type ReviewInstructorLeaveStrategy struct {
	Tx *gorm.DB
}

func NewReviewInstructorLeaveStrategy(tx *gorm.DB) *ReviewInstructorLeaveStrategy {
	return &ReviewInstructorLeaveStrategy{Tx: tx}
}

func (s *ReviewInstructorLeaveStrategy) Execute() error {
	controller := controller.NewLeaveInstructorHRController(s.Tx)
	return HandleReviewRequest(
		s.Tx,
		"instructor leave",
		controller.ReviewInstructorLeaveRequest,
	)
}
