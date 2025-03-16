package controller

import (
	"ModEd/asset/model"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SupplyController struct {
	Db *gorm.DB
}

func (c *SupplyController) GetAll() (*[]model.Supply, error) {
	supplies := new([]model.Supply)
	result := c.Db.Find(&supplies)
	return supplies, result.Error
}

func (c *SupplyController) GetByID(supplyID uuid.UUID) (*model.Supply, error) {
	supply := new(model.Supply)
	result := c.Db.First(&supply, "supply_id = ?", supplyID)
	return supply, result.Error
}

func (c *SupplyController) Create(body *model.Supply) error {
	result := c.Db.Create(body);
	return result.Error
}

func (c *SupplyController) Update(supplyID uuid.UUID, body *model.Supply) error {
	body.SupplyID = supplyID
	result := c.Db.Updates(body)
	return result.Error
}

func (c *SupplyController) Delete(supplyID uuid.UUID) error {
	result := c.Db.Model(&model.Supply{}).Where("supply_id = ?", supplyID).Update("deleted_at", time.Now())
	return result.Error
}
