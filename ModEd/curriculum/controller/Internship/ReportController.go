package controller

import (
	model "ModEd/curriculum/model/Internship"
	"fmt"

	"gorm.io/gorm"
)

type ReportController struct {
	DB *gorm.DB
}

func (rc *ReportController) UpdateReportScore(studentID string, Score int) error {
	var application model.InternshipApplication
	if err := rc.DB.Where("student_code = ?", studentID).First(&application).Error; err != nil {
		return fmt.Errorf("failed to find application for student_code '%s': %w", studentID, err)
	}
	var report model.InternshipReport
	if err := rc.DB.Where("id = ?", application.InternshipReportId).First(&report).Error; err != nil {
		return fmt.Errorf("failed to find report with id '%d': %w", application.InternshipReportId, err)
	}
	rc.DB.Model(&model.InternshipReport{}).Where("id = ?", report.ID).Update("ReportScore", Score)
	return nil
}
