package handler

import (
	"ModEd/hr/controller"

	"gorm.io/gorm"
)

type ReviewInstructorRaiseStrategy struct {
	Tx *gorm.DB
}

func NewReviewInstructorRaiseStrategy(tx *gorm.DB) *ReviewInstructorRaiseStrategy {
	return &ReviewInstructorRaiseStrategy{Tx: tx}
}

func (s *ReviewInstructorRaiseStrategy) Execute() error {
	controller := controller.CreateRaiseInstructorHRController(s.Tx)
	return HandleReviewRequest(
		s.Tx,
		"instructor raise",
		controller.ReviewInstructorRaiseRequest,
	)
}
