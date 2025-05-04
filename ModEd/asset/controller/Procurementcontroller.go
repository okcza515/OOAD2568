package controller

import (
	model "ModEd/asset/model"
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

func (c *ProcurementController) GetProcurementByID(id uint) (*model.Procurement, error) {
	var procurement model.Procurement
	err := c.db.First(&procurement, id).Error
	return &procurement, err
}

func (c *ProcurementController) Create(p *model.Procurement) error {
	return c.db.Create(p).Error
}

// Update procurement fields (generic)
func (c *ProcurementController) Update(id uint, updates map[string]any) error {
	return c.db.Model(&model.Procurement{}).
		Where("procurement_id = ?", id).
		Updates(updates).Error
}

// Approve procurement
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

// Reject procurement
func (c *ProcurementController) Reject(id uint, approverID uint) error {
	return c.db.Model(&model.Procurement{}).
		Where("procurement_id = ?", id).
		Updates(map[string]interface{}{
			"status":      model.ProcurementStatusRejected,
			"approver_id": approverID,
		}).Error
}

// Delete soft
func (c *ProcurementController) Delete(id uint) error {
	return c.db.Delete(&model.Procurement{}, id).Error
}

func (c *ProcurementController) OnApproved(id uint) error {
	return c.db.Model(&model.Procurement{}).
		Where("procurement_id = ?", id).
		Update("status", model.ProcurementStatusApproved).Error
}

func (c *ProcurementController) OnRejected(id uint) error {
	return c.db.Model(&model.Procurement{}).
		Where("procurement_id = ?", id).
		Update("status", model.ProcurementStatusRejected).Error
}
