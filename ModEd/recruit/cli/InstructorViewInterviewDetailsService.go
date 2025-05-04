// MEP-1003 Student Recruitment
package cli

import (
	"ModEd/recruit/controller"
	"ModEd/recruit/model"

	"gorm.io/gorm"
)

type InstructorViewInterviewDetailsService interface {
	ViewInterviewDetails(instructorID uint, status string) ([]model.Interview, error)
}

type instructorViewInterviewDetailsService struct {
	DB *gorm.DB
}

func NewInstructorViewInterviewDetailsService(DB *gorm.DB) InstructorViewInterviewDetailsService {
	return &instructorViewInterviewDetailsService{
		DB: DB,
	}
}

func (s *instructorViewInterviewDetailsService) ViewInterviewDetails(instructorID uint, statusfilter string) ([]model.Interview, error) {
	instructorCtrl := controller.NewInstructorController(s.DB)

	interviews, err := instructorCtrl.GetInterviewsByInstructor(instructorID)
	if err != nil {
		return nil, err
	}

	filtered := make([]model.Interview, 0)
	for _, interview := range interviews {
		if statusfilter == "all" || string(interview.InterviewStatus) == statusfilter {
			filtered = append(filtered, interview)
		}
	}

	return filtered, nil
}
