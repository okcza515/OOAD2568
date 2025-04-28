// MEP-1003 Student Recruitment
package cli

import (
	"ModEd/recruit/controller"

	"gorm.io/gorm"
)

type InstructorEvaluateApplicantService interface {
	EvaluateApplicant(interviewID uint, score float64) error
}

type instructorEvaluateApplicantService struct {
	DB *gorm.DB
}

func NewInstructorEvaluateApplicantService(DB *gorm.DB) *instructorEvaluateApplicantService {
	return &instructorEvaluateApplicantService{
		DB: DB,
	}
}

func (s *instructorEvaluateApplicantService) EvaluateApplicant(interviewID uint, score float64) error {
	instructorCtrl := controller.CreateInstructorController(s.DB)
	return instructorCtrl.EvaluateApplicant(interviewID, score)
}
