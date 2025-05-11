package handler

import (
	"ModEd/hr/controller"

	"gorm.io/gorm"
)

type ReviewStudentLeaveStrategy struct {
	Tx *gorm.DB
}

func NewReviewStudentLeaveStrategy(tx *gorm.DB) *ReviewStudentLeaveStrategy {
	return &ReviewStudentLeaveStrategy{Tx: tx}
}

func (s *ReviewStudentLeaveStrategy) Execute() error {
	controller := controller.NewLeaveStudentHRController(s.Tx)
	return HandleReviewRequest(
		s.Tx,
		"student leave",
		controller.ReviewStudentLeaveRequest,
	)
}
