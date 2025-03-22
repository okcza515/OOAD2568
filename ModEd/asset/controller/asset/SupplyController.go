package asset

import (
	"ModEd/asset/model/asset"
	"time"

	"gorm.io/gorm"
)

type SupplyController struct {
	db *gorm.DB
}

func (c *SupplyController) GetAll() (*[]asset.Supply, error) {
	supplies := new([]asset.Supply)
	result := c.db.Find(&supplies)
	return supplies, result.Error
}

func (c *SupplyController) GetByID(supplyID uint) (*asset.Supply, error) {
	supply := new(asset.Supply)
	result := c.db.First(&supply, "ID = ?", supplyID)
	return supply, result.Error
}

func (c *SupplyController) Create(body *asset.Supply) error {
	result := c.db.Create(body)
	return result.Error
}

func (c *SupplyController) Update(supplyID uint, body *asset.Supply) error {
	body.ID = supplyID
	result := c.db.Updates(body)
	return result.Error
}

func (c *SupplyController) Delete(supplyID uint) error {
	result := c.db.Model(&asset.Supply{}).Where("ID = ?", supplyID).Update("deleted_at", time.Now())
	return result.Error
}
