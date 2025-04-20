package controller

// MEP-1012 Asset

import (
	"ModEd/asset/model"
	"time"

	"gorm.io/gorm"
)

type SupplyLogController struct {
	db *gorm.DB
}

func (c *SupplyLogController) GetAll() (*[]model.SupplyLog, error) {
	suppliesLogs := new([]model.SupplyLog)
	result := c.db.Find(&suppliesLogs)
	return suppliesLogs, result.Error
}

func (c *SupplyLogController) GetByID(supplyLogID uint) (*model.SupplyLog, error) {
	supply := new(model.SupplyLog)
	result := c.db.First(&supply, "ID = ?", supplyLogID)
	return supply, result.Error
}

func (c *SupplyLogController) Create(body *model.SupplyLog) error {
	result := c.db.Create(body)
	return result.Error
}

func (c *SupplyLogController) Update(supplyLogID uint, body *model.SupplyLog) error {
	body.ID = supplyLogID
	result := c.db.Updates(body)
	return result.Error
}

func (c *SupplyLogController) Delete(supplyLogID uint) error {
	result := c.db.Model(&model.SupplyLog{}).Where("ID = ?", supplyLogID).Update("deleted_at", time.Now())
	return result.Error
}
