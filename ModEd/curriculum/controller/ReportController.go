//MEP-1009 Student Internship
package controller

import (
	"ModEd/core"
	model "ModEd/curriculum/model"
	"fmt"

	"gorm.io/gorm"
)

type ReportController struct {
	*core.BaseController
	Connector *gorm.DB
}

func CreateReportController(connector *gorm.DB) *ReportController {
	return &ReportController{
		Connector:      connector,
		BaseController: core.NewBaseController("Report", connector),
	}
}

func (rc *ReportController) UpdateReportScore(studentID string, Score int) error {
	var application model.InternshipApplication
	if err := rc.Connector.Where("student_code = ?", studentID).Last(&application).Error; err != nil {
		return fmt.Errorf("failed to find application for student_code '%s': %w", studentID, err)
	}
	var report model.InternshipReport
	if err := rc.Connector.Where("id = ?", application.InternshipReportId).Last(&report).Error; err != nil {
		return fmt.Errorf("failed to find report with id '%d': %w", application.InternshipReportId, err)
	}
	rc.Connector.Model(&model.InternshipReport{}).Where("id = ?", report.ID).Update("ReportScore", Score)
	return nil
}
