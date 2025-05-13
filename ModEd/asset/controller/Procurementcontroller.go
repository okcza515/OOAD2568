// MEP-1014
package controller

import (
	model "ModEd/asset/model"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type ProcurementController struct {
	db *gorm.DB
}

func CreateProcurementController(db *gorm.DB) *ProcurementController {
	return &ProcurementController{db: db}
}

func (c *ProcurementController) CreateProcurement(body *model.Procurement) error {
	return c.db.Create(body).Error
}

func (c *ProcurementController) ListAllProcurement() (*[]model.Procurement, error) {
	var procurements []model.Procurement
	err := c.db.Find(&procurements).Error
	return &procurements, err
}

func (c *ProcurementController) ListAllPendingProcurement() (*[]model.Procurement, error) {
	var procurements []model.Procurement
	err := c.db.
		Where("status = ?", "pending").
		Find(&procurements).Error
	return &procurements, err
}

func (c *ProcurementController) ListSelectedProcurements() (*[]model.Procurement, error) {
	var procurements []model.Procurement

	err := c.db.Preload("Approver").
		Preload("TOR").
		Joins("JOIN tors ON tors.tor_id = procurements.tor_id").
		Where("tors.status = ?", "selected").
		Find(&procurements).Error

	if err != nil {
		return nil, err
	}

	return &procurements, nil
}

func (c *ProcurementController) GetProcurementByID(id uint) (*model.Procurement, error) {
	var procurement model.Procurement
	err := c.db.First(&procurement, id).Error
	return &procurement, err
}

func (c *ProcurementController) Update(id uint, updates map[string]any) error {
	return c.db.Model(&model.Procurement{}).
		Where("procurement_id = ?", id).
		Updates(updates).Error
}

func (c *ProcurementController) Approve(id uint, approverID uint) error {
	now := time.Now()
	return c.db.Model(&model.Procurement{}).
		Where("procurement_id = ?", id).
		Updates(map[string]interface{}{
			"status":      model.ProcurementStatusApproved,
			"approver_id": approverID,
			"updated_at":  now,
		}).Error
}

func (c *ProcurementController) Reject(id uint, approverID uint) error {
	return c.db.Model(&model.Procurement{}).
		Where("procurement_id = ?", id).
		Updates(map[string]interface{}{
			"status":      model.ProcurementStatusRejected,
			"approver_id": approverID,
		}).Error
}

func (c *ProcurementController) DeleteProcurement(id uint) error {
	return c.db.Delete(&model.Procurement{}, id).Error
}

func (c *ProcurementController) OnApproved(id uint, approverID uint) error {
	return c.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&model.Procurement{}).
			Where("procurement_id = ?", id).
			Updates(map[string]interface{}{
				"status":        model.ProcurementStatusApproved,
				"approver_id":   approverID,
				"approval_time": time.Now(),
			}).Error; err != nil {
			return err
		}

		var existingApproval model.AcceptanceApproval
		err := tx.Where("procurement_id = ?", id).
			First(&existingApproval).Error

		if err == nil {
			if err := tx.Model(&existingApproval).
				Updates(map[string]interface{}{
					"status":        model.AcceptanceStatusPending,
					"approver_id":   approverID,
					"approval_time": time.Now(),
				}).Error; err != nil {
				return err
			}
			fmt.Printf("Updated existing Acceptance Approval with ID %d for Procurement ID %d\n", existingApproval.AcceptanceApprovalID, id)
		} else if err == gorm.ErrRecordNotFound {
			newApproval := model.AcceptanceApproval{
				ProcurementID: id,
				Status:        model.AcceptanceStatusPending,
				ApproverID:    &approverID,
				CreatedAt:     time.Now(),
			}

			if err := tx.Create(&newApproval).Error; err != nil {
				return err
			}
			fmt.Printf("Created new Acceptance Approval for Procurement ID %d with ID %d\n", id, newApproval.AcceptanceApprovalID)
		} else {
			return err
		}

		return nil
	})
}

func (c *ProcurementController) OnRejected(id uint, approverID uint) error {
	return c.db.Model(&model.Procurement{}).
		Where("procurement_id = ?", id).
		Updates(map[string]interface{}{
			"status":        model.ProcurementStatusRejected,
			"approver_id":   approverID,
			"approval_time": time.Now(),
		}).Error
}
