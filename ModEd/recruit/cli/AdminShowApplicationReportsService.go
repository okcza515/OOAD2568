package cli

import (
	"ModEd/recruit/controller"
	"ModEd/recruit/model"
)

type AdminShowApplicationReportsService interface {
	GetApplicationReport(applicantID uint) (*model.ApplicationReport, error)
}

type adminShowApplicationReportsService struct {
	reportCtrl *controller.ApplicationReportController
}

func NewAdminShowApplicationReportsService(ctrl *controller.ApplicationReportController) AdminShowApplicationReportsService {
	return &adminShowApplicationReportsService{
		reportCtrl: ctrl,
	}
}

func (s *adminShowApplicationReportsService) GetApplicationReport(applicantID uint) (*model.ApplicationReport, error) {
	return s.reportCtrl.GetApplicationReportByApplicantID(applicantID)
}
