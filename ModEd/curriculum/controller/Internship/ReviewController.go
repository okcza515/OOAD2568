package controller

import (
	"ModEd/core"
	model "ModEd/curriculum/model/Internship"
	"fmt"

	"gorm.io/gorm"
)

type ReviewController struct {
	*core.BaseController
	Connector *gorm.DB
}

func CreateReviewController(connector *gorm.DB) *ReviewController {
	return &ReviewController{
		Connector:      connector,
		BaseController: core.NewBaseController("Review", connector),
	}
}

func (rc *ReviewController) UpdateReviewScore(studentID string, SupervisorScore int, MentorScore int) error {
	var application model.InternshipApplication
	if err := rc.Connector.Where("student_code = ?", studentID).Last(&application).Error; err != nil {
		return fmt.Errorf("failed to find application for student_code '%s': %w", studentID, err)
	}
	var report model.SupervisorReview
	if err := rc.Connector.Where("id = ?", application.InternshipReportId).Last(&report).Error; err != nil {
		return fmt.Errorf("failed to find report with id '%d': %w", application.InternshipReportId, err)
	}
	if err := rc.Connector.Model(&model.SupervisorReview{}).Where("id = ?", report.ID).Updates(map[string]interface{}{
		"InstructorScore": SupervisorScore,
		"MentorScore":     MentorScore,
	}).Error; err != nil {
		return fmt.Errorf("failed to update review scores for report id '%d': %w", report.ID, err)
	}

	return nil
}
