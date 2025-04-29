package controller

import (
	"ModEd/core"
	"ModEd/project/model"

	"gorm.io/gorm"
)

type PresentationController struct {
	*core.BaseController[*model.Presentation]
	db *gorm.DB
}

func NewPresentationController(db *gorm.DB) *PresentationController {
	return &PresentationController{
		BaseController: core.NewBaseController[*model.Presentation](db),
		db:             db,
	}
}

func (c *PresentationController) ListAllPresentations() ([]model.Presentation, error) {
	var presentations []model.Presentation
	result := c.db.Find(&presentations)
	if result.Error != nil {
		return nil, result.Error
	}
	return presentations, nil
}

func (c *PresentationController) RetrievePresentation(id uint) (*model.Presentation, error) {
	var Presentation model.Presentation
	if err := c.db.Where("id = ?", id).First(&Presentation).Error; err != nil {
		return nil, err
	}
	return &Presentation, nil
}

func (c *PresentationController) InsertPresentation(Presentation *model.Presentation) error {
	return c.db.Create(Presentation).Error
}

func (c *PresentationController) UpdatePresentation(Presentation *model.Presentation) error {
	return c.db.Save(Presentation).Error
}

func (c *PresentationController) DeletePresentation(id uint) error {
	return c.db.Where("id = ?", id).Delete(&model.Presentation{}).Error
}
