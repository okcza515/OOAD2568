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
	procurement := ProcurementController{db: db}
	db.AutoMigrate(&model.Procurement{})
	return &procurement
}

// List all procurements with related data
func (c *ProcurementController) ListAll() ([]model.Procurement, error) {
	var procurements []model.Procurement
	err := c.db.
		Preload("TOR.InstrumentRequest.Instruments.Category").
		Preload("Approver").
		Find(&procurements).Error
	return procurements, err
}

// Get a procurement by ID
func (c *ProcurementController) GetByID(id uint) (*model.Procurement, error) {
	var procurement model.Procurement
	err := c.db.
		Preload("TOR.InstrumentRequest.Instruments.Category").
		Preload("Approver").
		First(&procurement, "procurement_id = ?", id).Error
	return &procurement, err
}

// Create a new procurement
func (c *ProcurementController) Create(p *model.Procurement) error {
	return c.db.Create(p).Error
}

// Update procurement fields
func (c *ProcurementController) Update(id uint, updates map[string]interface{}) error {
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
	now := time.Now()
	return c.db.Model(&model.Procurement{}).
		Where("procurement_id = ?", id).
		Updates(map[string]interface{}{
			"status":      model.ProcurementStatusRejected,
			"approver_id": approverID,
			"updated_at":  now,
		}).Error
}

// Soft delete procurement
func (c *ProcurementController) Delete(id uint) error {
	return c.db.Delete(&model.Procurement{}, id).Error
}
