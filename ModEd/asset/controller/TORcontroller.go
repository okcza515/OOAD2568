// MEP-1014
package controller

import (
	"time"

	model "ModEd/asset/model"

	"gorm.io/gorm"
)

type TORController struct {
	db *gorm.DB
}

func (c *TORController) CreateTOR(tor *model.TOR) error {
	return c.db.Create(tor).Error
}

func (c *TORController) GetAllTORs() (*[]model.TOR, error) {
	var tors []model.TOR
	err := c.db.Find(&tors).Error
	return &tors, err
}

func (c *TORController) GetTORByID(id uint) (*model.TOR, error) {
	var tor model.TOR
	err := c.db.
		Preload("InstrumentRequest.Instruments.Category").
		Preload("InstrumentRequest.Departments").
		First(&tor, "tor_id = ?", id).Error

	return &tor, err
}

func (c *TORController) DeleteTOR(id uint) error {
	return c.db.Model(&model.TOR{}).Where("tor_id = ?", id).Update("deleted_at", time.Now()).Error
}
