// MEP-1003 Student Recruitment
package controller

import (
	"ModEd/core"
	"ModEd/recruit/model"

	"gorm.io/gorm"
)

type ApplicationReportController struct {
	DB   *gorm.DB
	Base *core.BaseController[*model.ApplicationReport]
}

func NewApplicationReportController(db *gorm.DB) *ApplicationReportController {
	return &ApplicationReportController{
		Base: core.NewBaseController[*model.ApplicationReport](db),
		DB:   db,
	}
}

func (ctrl *ApplicationReportController) SaveApplicationReport(report *model.ApplicationReport) error {
	return ctrl.Base.Insert(report)
}

func (ctrl *ApplicationReportController) GetFilteredApplication(condition map[string]interface{}) ([]*model.ApplicationReport, error) {
	return ctrl.Base.List(
		condition,
		"Applicant",
		"ApplicationRound",
		"Faculty",
		"Department",
	)
}

func (ctrl *ApplicationReportController) UpdateApplicationStatus(applicantionreportID uint, newStatus model.ApplicationStatus) error {
	result := ctrl.DB.Model(&model.ApplicationReport{}).
		Where("application_report_id = ?", applicantionreportID).
		Update("application_statuses", newStatus)

	return result.Error
}

func (ctrl *ApplicationReportController) GetApplicationReportByID(reportID uint) ([]*model.ApplicationReport, error) {
	condition := map[string]interface{}{
		"application_report_id": reportID,
	}
	return ctrl.Base.List(
		condition,
		"Applicant",
		"ApplicationRound",
		"Faculty",
		"Department",
	)
}
