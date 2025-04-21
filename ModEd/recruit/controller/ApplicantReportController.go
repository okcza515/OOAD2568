// MEP-1003 Student Recruitment
package controller

import (
	"ModEd/core"
	"ModEd/recruit/model"

	"gorm.io/gorm"
)

type ApplicationReportController struct {
	DB   *gorm.DB
	Base *core.BaseController
}

func CreateApplicationReportController(db *gorm.DB) *ApplicationReportController {
	return &ApplicationReportController{
		Base: core.NewBaseController("ApplicationReport", db),
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

func (ctrl *ApplicationReportController) UpdateApplicationStatus(applicantID uint, newStatus model.ApplicationStatus) error {
	result := ctrl.DB.Model(&model.ApplicationReport{}).
		Where("applicant_id = ?", applicantID).
		Update("application_statuses", newStatus)

	return result.Error
}

func (ctrl *ApplicationReportController) GetApplicationStatusByID(applicantID uint) (model.ApplicationStatus, error) {
	var status model.ApplicationStatus
	err := ctrl.DB.
		Model(&model.ApplicationReport{}).
		Select("application_statuses").
		Where("applicant_id = ?", applicantID).
		Scan(&status).Error

	return status, err
}

func (ctrl *ApplicationReportController) GetFullApplicationReportByApplicantID(applicantID uint) (*model.ApplicationReport, error) {
	var report model.ApplicationReport
	err := ctrl.DB.Preload("Applicant").
		Preload("ApplicationRound").
		Preload("Faculty").
		Preload("Department").
		Where("applicant_id = ?", applicantID).
		First(&report).Error
	if err != nil {
		return nil, err
	}
	return &report, nil
}
