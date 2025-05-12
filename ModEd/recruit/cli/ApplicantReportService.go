// MEP-1003 Student Recruitment
package cli

import (
	"ModEd/recruit/controller"
	"ModEd/recruit/model"
	"errors"

	"gorm.io/gorm"
)

type ApplicantReportService interface {
	GetApplicationReport(applicantionReportID uint) (*model.ApplicationReport, error)
	DisplayReport([]*model.ApplicationReport)
	ConfirmAcceptance(applicationReportID uint, status model.ApplicationStatus) error
}

type applicantReportService struct {
	DB                  *gorm.DB
	ApplicantReportCtrl *controller.ApplicationReportController
}

func NewApplicantReportService(DB *gorm.DB, applicationReportCtrl *controller.ApplicationReportController) *applicantReportService {
	return &applicantReportService{
		DB:                  DB,
		ApplicantReportCtrl: applicationReportCtrl,
	}
}

func (s *applicantReportService) GetApplicationReport(applicantionReportID uint) (*model.ApplicationReport, error) {
	report := controller.ApplicationReport{
		Controller: s.ApplicantReportCtrl,
	}

	var condition map[string]interface{}
	condition = map[string]interface{}{
		"application_report_id": applicantionReportID,
	}

	filteredData, err := report.GetFilteredReport(condition)
	if err != nil {
		println("can't get report")
		return nil, err
	}
	if len(filteredData) == 0 {
		return nil, errors.New("no report found")
	}

	return filteredData[0], nil

}

func (s *applicantReportService) DisplayReport(reports []*model.ApplicationReport) {
	reportDisplay := controller.ApplicationReport{
		Controller: s.ApplicantReportCtrl,
	}
	converted := make([]model.ApplicationReport, len(reports))
	for i, r := range reports {
		converted[i] = *r
	}

	reportDisplay.DisplayReport(converted)
}

func (s *applicantReportService) ConfirmAcceptance(applicationReportID uint, status model.ApplicationStatus) error {
	err := s.ApplicantReportCtrl.UpdateApplicationStatus(applicationReportID, status)
	if err != nil {
		return err
	}
	return nil
}
