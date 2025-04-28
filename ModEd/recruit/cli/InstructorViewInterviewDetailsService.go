// MEP-1003 Student Recruitment
package cli

import (
	"ModEd/recruit/controller"
	"ModEd/recruit/model"

	"gorm.io/gorm"
)

type InstructorViewInterviewDetailsService interface {
	ViewInterviewDetails(instructorID uint) ([]model.Interview, error)
}

type instructorViewInterviewDetailsService struct {
	DB *gorm.DB
}

func NewInstructorViewInterviewDetailsService(DB *gorm.DB) InstructorViewInterviewDetailsService {
	return &instructorViewInterviewDetailsService{
		DB: DB,
	}
}

func (s *instructorViewInterviewDetailsService) ViewInterviewDetails(instructorID uint) ([]model.Interview, error) {
	instructorCtrl := controller.CreateInstructorController(s.DB)
	return instructorCtrl.GetInterviewsByInstructor(instructorID)
}
