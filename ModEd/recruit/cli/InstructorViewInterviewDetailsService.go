// MEP-1003 Student Recruitment
package cli

import (
	"ModEd/recruit/controller"
	"ModEd/recruit/model"

	"gorm.io/gorm"
)

type InstructorViewInterviewDetailsService interface {
	ViewInterviewDetails(instructorID uint, status string, interviewCtrl *controller.InterviewController) ([]model.Interview, error)
}

type instructorViewInterviewDetailsService struct {
	DB *gorm.DB
}

func NewInstructorViewInterviewDetailsService(DB *gorm.DB) InstructorViewInterviewDetailsService {
	return &instructorViewInterviewDetailsService{
		DB: DB,
	}
}

func (s *instructorViewInterviewDetailsService) ViewInterviewDetails(
	instructorID uint,
	status string,
	interviewCtrl *controller.InterviewController,
) ([]model.Interview, error) {

	filters := []controller.FilterStrategy{
		&controller.FilterByInstructorID{InstructorID: instructorID},
		&controller.FilterByStatus{Status: status},
	}

	report := controller.InterviewReport{
		InterviewProvider: interviewCtrl,
		Filters:           filters,
	}

	rawData, err := report.GetReport()
	if err != nil {
		println("can't get report")
		return nil, err
	}

	filteredData, err := report.FilterReport(rawData)
	if err != nil {
		println("can't filter report")
		return nil, err
	}
	return filteredData, nil
}
