package handler

import (
	"ModEd/hr/controller"

	"gorm.io/gorm"
)

type ReviewInstructorResignStrategy struct {
	Tx *gorm.DB
}

func NewReviewInstructorResignStrategy(tx *gorm.DB) *ReviewInstructorResignStrategy {
	return &ReviewInstructorResignStrategy{Tx: tx}
}

func (s *ReviewInstructorResignStrategy) Execute() error {
	controller := controller.NewResignationInstructorHRController(s.Tx)
	return HandleReviewRequest(
		s.Tx,
		"instructor resign",
		controller.ReviewInstructorResignRequest,
	)
}
