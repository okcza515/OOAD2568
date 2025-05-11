package handler

import (
	"ModEd/hr/controller"

	"gorm.io/gorm"
)

type ReviewStudentResignStrategy struct {
	Tx *gorm.DB
}

func NewReviewStudentResignStrategy(tx *gorm.DB) *ReviewStudentResignStrategy {
	return &ReviewStudentResignStrategy{Tx: tx}
}

func (s *ReviewStudentResignStrategy) Execute() error {
	controller := controller.NewResignationStudentHRController(s.Tx)
	return HandleReviewRequest(
		s.Tx,
		"student resign",
		controller.ReviewStudentResignRequest,
	)
}
