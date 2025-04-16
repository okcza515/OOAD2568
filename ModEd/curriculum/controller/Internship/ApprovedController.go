package controller

import (
    model "ModEd/curriculum/model/Internship"
    "fmt"
    "gorm.io/gorm"
)

type ApprovedController struct {
    Db *gorm.DB
}

func NewApprovedController(db *gorm.DB) *ApprovedController {
    return &ApprovedController{Db: db}
}

func (c *ApprovedController) UpdateAdvisorApprovalStatus(applicationID uint, status model.ApprovedStatus) error {
    if status != model.APPROVED && status != model.REJECT {
        return fmt.Errorf("invalid status: %s", status)
    }

    result := c.Db.Model(&model.InternshipApplication{}).
        Where("id = ?", applicationID).
        Update("approval_advisor_status", status)

    if result.Error != nil {
        return fmt.Errorf("failed to update advisor approval status: %w", result.Error)
    }

    return nil
}

func (c *ApprovedController) UpdateCompanyApprovalStatus(applicationID uint, status model.ApprovedStatus) error {
    if status != model.APPROVED && status != model.REJECT {
        return fmt.Errorf("invalid status: %s", status)
    }

    result := c.Db.Model(&model.InternshipApplication{}).
        Where("id = ?", applicationID).
        Update("approval_company_status", status)

    if result.Error != nil {
        return fmt.Errorf("failed to update company approval status: %w", result.Error)
    }

    return nil
}