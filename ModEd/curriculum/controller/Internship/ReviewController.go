package controller

import (
	model "ModEd/curriculum/model/Internship"
	"fmt"

	"gorm.io/gorm"
)

type ReviewController struct {
	DB *gorm.DB
}

func (rc *ReviewController) UpdateReviewScore(studentID string, SupervisorScore int, MentorScore int) error {
	var application model.InternshipApplication
	if err := rc.DB.Where("student_code = ?", studentID).First(&application).Error; err != nil {
		return fmt.Errorf("failed to find application for student_code '%s': %w", studentID, err)
	}
	var report model.SupervisorReview
	if err := rc.DB.Where("id = ?", application.InternshipReportId).First(&report).Error; err != nil {
		return fmt.Errorf("failed to find report with id '%d': %w", application.InternshipReportId, err)
	}
	if err := rc.DB.Model(&model.SupervisorReview{}).Where("id = ?", report.ID).Updates(map[string]interface{}{
		"InstructorScore": SupervisorScore,
		"MentorScore":     MentorScore,
	}).Error; err != nil {
		return fmt.Errorf("failed to update review scores for report id '%d': %w", report.ID, err)
	}

	return nil
}
