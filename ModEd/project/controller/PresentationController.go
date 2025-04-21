package controller

import (
	"ModEd/project/model"

	"gorm.io/gorm"
)

type IPresentationController interface {
	ListAllPresentations() ([]model.Presentation, error)
	RetrievePresentation(id uint) (*model.Presentation, error)
	InsertPresentation(Presentation *model.Presentation) error
	UpdatePresentation(Presentation *model.Presentation) error
	DeletePresentation(id uint) error
}

type PresentationController struct {
	db *gorm.DB
}

func NewPresentationController(db *gorm.DB) IPresentationController {
	return &PresentationController{db: db}
}

func (c *PresentationController) ListAllPresentations() ([]model.Presentation, error) {
	var Presentations []model.Presentation
	err := c.db.Find(&Presentations).Error
	return Presentations, err
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
