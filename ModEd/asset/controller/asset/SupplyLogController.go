package asset

import (
	"ModEd/asset/model/asset"
	"time"

	"gorm.io/gorm"
)

type SupplyLogController struct {
	Db *gorm.DB
}

func (c *SupplyLogController) GetAll() (*[]asset.SupplyLog, error) {
	suppliesLogs := new([]asset.SupplyLog)
	result := c.Db.Find(&suppliesLogs)
	return suppliesLogs, result.Error
}

func (c *SupplyLogController) GetByID(supplyLogID uint) (*asset.SupplyLog, error) {
	supply := new(asset.SupplyLog)
	result := c.Db.First(&supply, "ID = ?", supplyLogID)
	return supply, result.Error
}

func (c *SupplyLogController) Create(body *asset.SupplyLog) error {
	result := c.Db.Create(body)
	return result.Error
}

func (c *SupplyLogController) Update(supplyLogID uint, body *asset.SupplyLog) error {
	body.ID = supplyLogID
	result := c.Db.Updates(body)
	return result.Error
}

func (c *SupplyLogController) Delete(supplyLogID uint) error {
	result := c.Db.Model(&asset.SupplyLog{}).Where("ID = ?", supplyLogID).Update("deleted_at", time.Now())
	return result.Error
}
