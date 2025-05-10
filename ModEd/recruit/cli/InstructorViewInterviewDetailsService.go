// MEP-1003 Student Recruitment
package cli

import (
	"ModEd/recruit/controller"
	"ModEd/recruit/model"

	"gorm.io/gorm"
)

type InstructorViewInterviewDetailsService interface {
	ViewInterviewDetails(instructorID uint, status string) ([]*model.Interview, error)
	DisplayReport([]*model.Interview)
}

type instructorViewInterviewDetailsService struct {
	DB            *gorm.DB
	InterviewCtrl *controller.InterviewController
}

func NewInstructorViewInterviewDetailsService(DB *gorm.DB, interviewCtrl *controller.InterviewController) InstructorViewInterviewDetailsService {
	return &instructorViewInterviewDetailsService{
		DB:            DB,
		InterviewCtrl: interviewCtrl,
	}
}

func (s *instructorViewInterviewDetailsService) ViewInterviewDetails(
	instructorID uint,
	status string,
) ([]*model.Interview, error) {

	report := controller.InterviewReport{
		Controller: s.InterviewCtrl,
	}

	var condition map[string]interface{}
	if status == "all" {
		condition = map[string]interface{}{
			"instructor_id": instructorID,
		}
	} else {
		condition = map[string]interface{}{
			"instructor_id":    instructorID,
			"interview_status": status,
		}
	}

	filteredData, err := report.GetFilteredInterviews(condition)
	if err != nil {
		println("can't get report")
		return nil, err
	}
	return filteredData, nil
}

func (s *instructorViewInterviewDetailsService) DisplayReport(interviews []*model.Interview) {
	reportDisplay := controller.InterviewReport{
		Controller: s.InterviewCtrl,
	}

	reportDisplay.DisplayReport(interviews)
}
