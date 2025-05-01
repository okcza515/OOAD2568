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

func CreateApplicationReportController(db *gorm.DB) *ApplicationReportController {
	return &ApplicationReportController{
		Base: core.NewBaseController[*model.ApplicationReport](db),
		DB:   db,
	}
}

func (ctrl *ApplicationReportController) SaveApplicationReport(report *model.ApplicationReport) error {
	return ctrl.Base.Insert(report)
}

func (ctrl *ApplicationReportController) GetApplicationReportByApplicantID(applicantID uint) (*model.ApplicationReport, error) {
	var report model.ApplicationReport
	err := ctrl.DB.Where("applicant_id = ?", applicantID).First(&report).Error
	if err != nil {
		return nil, err
	}
	return &report, nil
}

func (ctrl *ApplicationReportController) UpdateApplicationStatus(applicantionreportID uint, newStatus model.ApplicationStatus) error {
	result := ctrl.DB.Model(&model.ApplicationReport{}).
		Where("application_report_id = ?", applicantionreportID).
		Update("application_statuses", newStatus)

	return result.Error
}

func (ctrl *ApplicationReportController) GetApplicationReportByID(reportID uint) (*model.ApplicationReport, error) {
	var report model.ApplicationReport
	if err := ctrl.DB.Preload("Applicant").First(&report, reportID).Error; err != nil {
		return nil, err
	}
	return &report, nil
}

func (ctrl *ApplicationReportController) GetApplicationStatusByApplicantID(applicantID uint) (model.ApplicationStatus, error) {
	var status model.ApplicationStatus
	err := ctrl.DB.
		Model(&model.ApplicationReport{}).
		Select("application_statuses").
		Where("applicant_id = ?", applicantID).
		Scan(&status).Error

	return status, err
}

func (ctrl *ApplicationReportController) GetFullApplicationReportByApplicationID(applicantionReportID uint) (*model.ApplicationReport, error) {
	var report model.ApplicationReport
	err := ctrl.DB.Preload("Applicant").
		Preload("ApplicationRound").
		Preload("Faculty").
		Preload("Department").
		Where("application_report_id = ?", applicantionReportID).
		First(&report).Error
	if err != nil {
		return nil, err
	}
	return &report, nil
}
