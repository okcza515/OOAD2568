package controller

import (
	model "ModEd/asset/model"
	"gorm.io/gorm"
)

type TORController struct {
	db *gorm.DB
}

func CreateTORController(db *gorm.DB) *TORController {
	return &TORController{db: db}
}

func (c *TORController) CreateTOR(tor *model.TOR) error {
	return c.db.Create(tor).Error
}

func (c *TORController) GetAllTORs() ([]model.TOR, error) {
	var tors []model.TOR
	err := c.db.
		Preload("InstrumentRequest.Instruments.Category").
		Find(&tors).Error
	return tors, err
}

func (c *TORController) GetTORByID(id uint) (*model.TOR, error) {
	var tor model.TOR
	err := c.db.
		Preload("InstrumentRequest.Instruments.Category").
		First(&tor, "tor_id = ?", id).Error
	return &tor, err
}

func (c *TORController) DeleteTOR(id uint) error {
	return c.db.Delete(&model.TOR{}, id).Error
}
