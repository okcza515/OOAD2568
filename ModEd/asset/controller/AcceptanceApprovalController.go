// MEP-1014
package controller

import (
	model "ModEd/asset/model"
	"errors"
	"gorm.io/gorm"
)

type AcceptanceApprovalController struct {
	db *gorm.DB
}

func (c *AcceptanceApprovalController) CreateAcceptanceApproval(body *model.AcceptanceApproval) error {
	return c.db.Create(body).Error
}

func (c *AcceptanceApprovalController) ListAllAcceptanceApprovals() (*[]model.AcceptanceApproval, error) {
	var approvals  []model.AcceptanceApproval
	err := c.db.Preload("Procurement").
		Preload("Approver").
		Find(&approvals ).Error
	return &approvals , err
}

func (c *AcceptanceApprovalController) GetAcceptanceApprovalByID(id uint) (*model.AcceptanceApproval, error) {
	var approvals  model.AcceptanceApproval
	err := c.db.
		Preload("Procurement").
		Preload("Approver").
		First(&approvals , id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &approvals , err
}

func (c *AcceptanceApprovalController) UpdateAcceptanceApprovalStatus(id uint, status model.AcceptanceStatus, approverID uint) error {
	return c.db.Model(&model.AcceptanceApproval{}).
		Where("acceptance_approval_id = ?", id).
		Updates(map[string]interface{}{
			"status":      status,
			"approver_id": approverID,
		}).Error
}

func (c *AcceptanceApprovalController) DeleteAcceptanceApproval(id uint) error {
	return c.db.Model(&model.AcceptanceApproval{}).
		Where("acceptance_approval_id = ?", id).
		Delete(&model.AcceptanceApproval{}).Error
}

func (c *AcceptanceApprovalController) GetQuotationDetailsByProcurement(procurementID uint) ([]model.QuotationDetail, error) {
	var quotations []model.Quotation

	err := c.db.Preload("Details").
		Where("procurement_id = ?", procurementID).
		Find(&quotations).Error
	if err != nil {
		return nil, err
	}

	var details []model.QuotationDetail
	for _, quotation := range quotations {
		details = append(details, quotation.Details...)
	}

	return details, nil
}

func (c *AcceptanceApprovalController) GetCategoriesByIDs(ids []uint) ([]model.Category, error) {
	var categories []model.Category
	if len(ids) == 0 {
		return categories, nil
	}

	err := c.db.Where("id IN ?", ids).Find(&categories).Error
	return categories, err
}

func (c *AcceptanceApprovalController) OnApproved(id uint, approverID uint) error {
	return c.db.Transaction(func(tx *gorm.DB) error {		
		if err := tx.Model(&model.AcceptanceApproval{}).
			Where("acceptance_approval_id = ?", id).
			Updates(map[string]interface{}{
				"status":      model.AcceptanceStatusApproved,
				"approver_id": approverID,
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
				"status":      model.AcceptanceStatusRejected,
				"approver_id": approverID,
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