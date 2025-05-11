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

func (c *BudgetApprovalController) ListAllApprovals() ([]model.BudgetApproval, error) {
	var approvals []model.BudgetApproval
	err := c.db.
		Preload("Approver").
		// Preload("InstrumentRequest.Departments").
		Preload("InstrumentRequest").
		Find(&approvals).Error
	return approvals, err
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

func (c *BudgetApprovalController) OnApproved(id uint, approverID uint) error {
	return c.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&model.BudgetApproval{}).
			Where("budget_approval_id = ?", id).
			Updates(map[string]interface{}{
				"status":      model.BudgetStatusApproved,
				"approver_id": approverID,
				"approval_time": time.Now(),
			}).Error; err != nil {
			return err
		}

		var approval model.BudgetApproval
		if err := tx.First(&approval, id).Error; err != nil {
			return err
		}

		if err := tx.Model(&model.InstrumentRequest{}).
			Where("instrument_request_id = ?", approval.InstrumentRequestID).
			Update("status", model.InstrumentRequestStatusApproved).Error; err != nil {
			return err
		}

		return nil
	})
}

func (c *BudgetApprovalController) OnRejected(id uint, approverID uint) error {
	return c.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&model.BudgetApproval{}).
			Where("budget_approval_id = ?", id).
			Updates(map[string]interface{}{
				"status":      model.BudgetStatusRejected,
				"approver_id": approverID,
				"approval_time": time.Now(),
			}).Error; err != nil {
			return err
		}

		var approval model.BudgetApproval
		if err := tx.First(&approval, id).Error; err != nil {
			return err
		}

		if err := tx.Model(&model.InstrumentRequest{}).
			Where("instrument_request_id = ?", approval.InstrumentRequestID).
			Update("status", model.InstrumentRequestStatusRejected).Error; err != nil {
			return err
		}

		return nil
	})
}
