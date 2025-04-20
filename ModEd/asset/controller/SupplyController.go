package controller

// MEP-1012 Asset

import (
	"ModEd/asset/model"
	"time"

	"gorm.io/gorm"
)

type SupplyController struct {
	db *gorm.DB
}

func (c *SupplyController) GetAll() (*[]model.Supply, error) {
	supplies := new([]model.Supply)
	result := c.db.Find(&supplies)
	return supplies, result.Error
}

func (c *SupplyController) GetByID(supplyID uint) (*model.Supply, error) {
	supply := new(model.Supply)
	result := c.db.First(&supply, "ID = ?", supplyID)
	return supply, result.Error
}

func (c *SupplyController) Create(body *model.Supply) error {
	result := c.db.Create(body)
	return result.Error
}

func (c *SupplyController) Update(supplyID uint, body *model.Supply) error {
	body.ID = supplyID
	result := c.db.Updates(body)
	return result.Error
}

func (c *SupplyController) Delete(supplyID uint) error {
	result := c.db.Model(&model.Supply{}).Where("ID = ?", supplyID).Update("deleted_at", time.Now())
	return result.Error
}
