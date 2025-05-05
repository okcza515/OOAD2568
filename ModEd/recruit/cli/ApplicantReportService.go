package cli

import (
	"ModEd/recruit/controller"
	"ModEd/recruit/model"
	"fmt"

	"gorm.io/gorm"
)

type ApplicantReportService interface {
	GetFullApplicationReportByApplicationID(applicantionReportID uint) ([]model.ApplicationReport, error)
}

type InterviewService interface {
	GetInterviewDetails(applicantID uint) (*model.Interview, error)
}

type applicantReportService struct {
	DB *gorm.DB
}

type interviewService struct {
	DB *gorm.DB
}

func NewApplicantReportService(DB *gorm.DB) *applicantReportService {
	return &applicantReportService{
		DB: DB,
	}
}

func (s *applicantReportService) GetFullApplicationReportByApplicationID(applicantionReportID uint) ([]model.ApplicationReport, error) {
	provider := controller.NewApplicationReportController(s.DB)

	report := controller.ApplicationReport{
		ApplicationProvider: provider,
		Filters: []controller.FilterStrategy[model.ApplicationReport]{
			&controller.ApplicationReportFilterByID{ApplicationReportID: applicantionReportID},
		},
	}

	allData, err := report.GetReport()
	if err != nil {
		return nil, err
	}

	filtered, err := report.FilterReport(allData)
	if err != nil {
		return nil, err
	}

	if len(filtered) == 0 {
		return nil, fmt.Errorf("no report found with ID %d", applicantionReportID)
	}

	return filtered, nil
}

func NewInterviewService(DB *gorm.DB) *interviewService {
	return &interviewService{
		DB: DB,
	}
}

func (s *interviewService) GetInterviewDetails(applicantID uint) (*model.Interview, error) {
	return controller.GetInterviewDetails(s.DB, applicantID)
}
