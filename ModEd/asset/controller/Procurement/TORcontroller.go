// MEP-1014
package procurement

import (
	model "ModEd/asset/model/Procurement"
	"time"

	"gorm.io/gorm"
)

type TORController struct {
	db *gorm.DB
}

// CreateTOR creates a new TOR linked to an approved ItemRequest
func (c *TORController) CreateTOR(tor *model.TOR) error {
	return c.db.Create(tor).Error
}

// GetAllTORs retrieves all TOR records
func (c *TORController) GetAllTORs() (*[]model.TOR, error) {
	var tors []model.TOR
	err := c.db.Find(&tors).Error
	return &tors, err
}

// GetTORByID retrieves a TOR by its ID, including its related ItemRequest
func (c *TORController) GetTORByID(id uint) (*model.TOR, error) {
	tor := new(model.TOR)
	err := c.db.Preload("ItemRequest").First(&tor, "tor_id = ?", id).Error
	return tor, err
}

// DeleteTOR performs a soft delete on the TOR
func (c *TORController) DeleteTOR(id uint) error {
	return c.db.Model(&model.TOR{}).Where("tor_id = ?", id).Update("deleted_at", time.Now()).Error
}
