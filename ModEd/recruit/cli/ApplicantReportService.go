package cli

import (
	"ModEd/recruit/controller"
	"ModEd/recruit/model"

	"gorm.io/gorm"
)

type ApplicantReportService interface {
	GetFullApplicationReportByApplicationID(applicantionReportID uint) (*model.ApplicationReport, error)
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

func (s *applicantReportService) GetFullApplicationReportByApplicationID(applicantionReportID uint) (*model.ApplicationReport, error) {
	reportController := controller.NewApplicationReportController(s.DB)
	return reportController.GetFullApplicationReportByApplicationID(applicantionReportID)
}

func NewInterviewService(DB *gorm.DB) *interviewService {
	return &interviewService{
		DB: DB,
	}
}

func (s *interviewService) GetInterviewDetails(applicantID uint) (*model.Interview, error) {
	return controller.GetInterviewDetails(s.DB, applicantID)
}
