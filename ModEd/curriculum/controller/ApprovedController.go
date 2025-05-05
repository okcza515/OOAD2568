// MEP-1009 Student Internship
package controller

import (
	"ModEd/core"
	model "ModEd/curriculum/model"
	"fmt"

	"gorm.io/gorm"
)

type ApprovedController struct {
	*core.BaseController[model.InternshipApplication]
	Connector *gorm.DB
}

func NewApprovedController(connector *gorm.DB) *ApprovedController {
	return &ApprovedController{
		Connector:      connector,
		BaseController: core.NewBaseController[model.InternshipApplication](connector),
	}
}

func isValidStatus(status model.ApprovedStatus) bool {
	return status == model.APPROVED || status == model.REJECT
}

func (c *ApprovedController) UpdateAdvisorApprovalStatus(applicationID uint, status model.ApprovedStatus) error {
	if !isValidStatus(status) {
		return fmt.Errorf("invalid status: %s", status)
	}

	result := c.Connector.Model(&model.InternshipApplication{}).
		Where("id = ?", applicationID).
		Update("approval_advisor_status", status)

	if result.Error != nil {
		return fmt.Errorf("failed to update advisor approval status: %w", result.Error)
	}

	return nil
}

func (c *ApprovedController) UpdateCompanyApprovalStatus(applicationID uint, status model.ApprovedStatus) error {
	if !isValidStatus(status) {
		return fmt.Errorf("invalid status: %s", status)
	}

	result := c.Connector.Model(&model.InternshipApplication{}).
		Where("id = ?", applicationID).
		Update("approval_company_status", status)

	if result.Error != nil {
		return fmt.Errorf("failed to update company approval status: %w", result.Error)
	}
	return nil
}

func (ac *ApprovedController) UpdateApprovalStatuses(studentCode string, advisorStatus model.ApprovedStatus, companyStatus model.ApprovedStatus) error {
	var application model.InternshipApplication

	if err := ac.Connector.Where("student_code = ?", studentCode).Last(&application).Error; err != nil {
		return fmt.Errorf("internship application for student '%s' not found", studentCode)
	}

	application.ApprovalAdvisorStatus = advisorStatus
	application.ApprovalCompanyStatus = companyStatus

	if err := ac.Connector.Save(&application).Error; err != nil {
		return fmt.Errorf("failed to update application statuses: %w", err)
	}

	if advisorStatus == model.APPROVED && companyStatus == model.APPROVED {
		var student model.InternStudent
		if err := ac.Connector.Where("student_code = ?", studentCode).First(&student).Error; err != nil {
			return fmt.Errorf("failed to find student: %w", err)
		}

		student.InternStatus = model.ACTIVES

		if err := ac.Connector.Save(&student).Error; err != nil {
			return fmt.Errorf("failed to update student intern status: %w", err)
		}
	}

	return nil
}
