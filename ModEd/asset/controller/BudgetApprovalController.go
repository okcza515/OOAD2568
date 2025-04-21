// MEP-1014
package controller

import (
	"time"

	model "ModEd/asset/model"

	"gorm.io/gorm"
)

type BudgetApprovalController struct {
	db *gorm.DB
}

func (c *BudgetApprovalController) CreateBudgetRequest(req *model.BudgetApproval) error {
	return c.db.Create(req).Error
}

func (c *BudgetApprovalController) BudgetApprove(id uint, status model.BudgetApprovalStatus) error {
	return c.db.Model(&model.BudgetApproval{}).Where("budget_approval_id = ?", id).Update("status", status).Error
}

func (c *BudgetApprovalController) ShowBudgetRequestList(instrumentRequestID uint) (*[]model.BudgetApproval, error) {
	var approvals []model.BudgetApproval
	err := c.db.
		Preload("Approver").
		Preload("InstrumentRequest.Departments").
		Where("instrument_request_id = ?", instrumentRequestID).
		Find(&approvals).Error
	return &approvals, err
}

func (c *BudgetApprovalController) ShowBudgetRequestStatus(id uint) (*model.BudgetApproval, error) {
	approval := new(model.BudgetApproval)
	err := c.db.
		Preload("Approver").
		Preload("InstrumentRequest.Departments").
		First(&approval, "budget_approval_id = ?", id).Error
	return approval, err
}

func (c *BudgetApprovalController) ShowBudgetRequestByStatus(status model.BudgetApprovalStatus) (*[]model.BudgetApproval, error) {
	var approvals []model.BudgetApproval
	err := c.db.
		Preload("Approver").
		Preload("InstrumentRequest").
		Where("status = ?", status).
		Find(&approvals).Error
	return &approvals, err
}

func (c *BudgetApprovalController) DeleteBudgetRequest(id uint) error {
	return c.db.Model(&model.BudgetApproval{}).Where("budget_approval_id = ?", id).Update("deleted_at", time.Now()).Error
}
