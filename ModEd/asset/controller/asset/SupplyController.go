package asset

import (
	"ModEd/asset/model/asset"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SupplyController struct {
	Db *gorm.DB
}

func (c *SupplyController) GetAll() (*[]asset.Supply, error) {
	supplies := new([]asset.Supply)
	result := c.Db.Find(&supplies)
	return supplies, result.Error
}

func (c *SupplyController) GetByID(supplyID uuid.UUID) (*asset.Supply, error) {
	supply := new(asset.Supply)
	result := c.Db.First(&supply, "supply_id = ?", supplyID)
	return supply, result.Error
}

func (c *SupplyController) Create(body *asset.Supply) error {
	result := c.Db.Create(body)
	return result.Error
}

func (c *SupplyController) Update(supplyID uuid.UUID, body *asset.Supply) error {
	body.SupplyID = supplyID
	result := c.Db.Updates(body)
	return result.Error
}

func (c *SupplyController) Delete(supplyID uuid.UUID) error {
	result := c.Db.Model(&asset.Supply{}).Where("supply_id = ?", supplyID).Update("deleted_at", time.Now())
	return result.Error
}
