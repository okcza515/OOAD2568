// MEP-1014
package controller

import (
	"time"
	model "ModEd/asset/model"
	"gorm.io/gorm"
)

type AcceptanceApprovalController struct {
	db *gorm.DB
}

func (c *AcceptanceApprovalController) CreateAcceptanceRequest(req *model.AcceptanceApproval) error {
	return c.db.Create(req).Error
}

func (c *AcceptanceApprovalController) ListAllApprovals() ([]model.AcceptanceApproval, error) {
	var approvals []model.AcceptanceApproval
	err := c.db.
		Preload("Procurement").
		Preload("Approver").
		Find(&approvals).Error
	return approvals, err
}

func (c *AcceptanceApprovalController) ShowAcceptanceRequestList(procurementID uint) ([]model.AcceptanceApproval, error) {
	var approvals []model.AcceptanceApproval
	err := c.db.
		Preload("Procurement").
		Preload("Approver").
		Where("procurement_id = ?", procurementID).
		Find(&approvals).Error
	return approvals, err
}

func (c *AcceptanceApprovalController) ShowAcceptanceRequestStatus(id uint) (*model.AcceptanceApproval, error) {
	approval := new(model.AcceptanceApproval)
	err := c.db.
		Preload("Procurement").
		Preload("Approver").
		First(approval, id).Error
	return approval, err
}

func (c *AcceptanceApprovalController) ShowAcceptanceRequestByStatus(status model.AcceptanceStatus) ([]model.AcceptanceApproval, error) {
	var approvals []model.AcceptanceApproval
	err := c.db.
		Preload("Procurement").
		Preload("Approver").
		Where("status = ?", status).
		Find(&approvals).Error
	return approvals, err
}

func (c *AcceptanceApprovalController) DeleteAcceptanceRequest(id uint) error {
	return c.db.Model(&model.AcceptanceApproval{}).
		Where("acceptance_approval_id = ?", id).
		Update("deleted_at", time.Now()).Error
}

func (c *AcceptanceApprovalController) OnApproved(id uint, approverID uint) error {
	return c.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&model.AcceptanceApproval{}).
			Where("acceptance_approval_id = ?", id).
			Updates(map[string]interface{}{
				"status":       model.AcceptanceStatusApproved,
				"approver_id":  approverID,
				"approval_time": time.Now(),
			}).Error; err != nil {
			return err
		}

		var approval model.AcceptanceApproval
		if err := tx.First(&approval, id).Error; err != nil {
			return err
		}

		if err := tx.Model(&model.Procurement{}).
			Where("procurement_id = ?", approval.ProcurementID).
			Update("status", model.ProcurementStatusApproved).Error; err != nil {
			return err
		}

		return nil
	})
}

func (c *AcceptanceApprovalController) OnRejected(id uint, approverID uint) error {
	return c.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&model.AcceptanceApproval{}).
			Where("acceptance_approval_id = ?", id).
			Updates(map[string]interface{}{
				"status":       model.AcceptanceStatusRejected,
				"approver_id":  approverID,
				"approval_time": time.Now(),
			}).Error; err != nil {
			return err
		}

		var approval model.AcceptanceApproval
		if err := tx.First(&approval, id).Error; err != nil {
			return err
		}

		if err := tx.Model(&model.Procurement{}).
			Where("procurement_id = ?", approval.ProcurementID).
			Update("status", model.ProcurementStatusRejected).Error; err != nil {
			return err
		}

		return nil
	})
}
