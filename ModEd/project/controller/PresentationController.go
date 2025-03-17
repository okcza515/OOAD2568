package controller

import (
	"ModEd/project/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IPresentationController interface {
	ListAllPresentations() ([]model.Presentation, error)
	RetrievePresentation(id uuid.UUID) (*model.Presentation, error)
	InsertPresentation(presentation *model.Presentation) error
	UpdatePresentation(presentation *model.Presentation) error
	DeletePresentation(id uuid.UUID) error
}

type PresentationController struct {
	db *gorm.DB
}

func NewPresentationController(db *gorm.DB) IPresentationController {
	return &PresentationController{}
}

func (c *PresentationController) ListAllPresentations() ([]model.Presentation, error) {
	var presentations []model.Presentation
	err := c.db.Find(&presentations).Error
	return presentations, err
}

func (c *PresentationController) RetrievePresentation(id uuid.UUID) (*model.Presentation, error) {
	var presentation model.Presentation
	if err := c.db.Where("id = ?", id).First(&presentation).Error; err != nil {
		return nil, err
	}
	return &presentation, nil
}

func (c *PresentationController) InsertPresentation(presentation *model.Presentation) error {
	return c.db.Create(presentation).Error
}

func (c *PresentationController) UpdatePresentation(presentation *model.Presentation) error {
	return c.db.Save(presentation).Error
}

func (c *PresentationController) DeletePresentation(id uuid.UUID) error {
	return c.db.Where("id = ?", id).Delete(&model.Presentation{}).Error
}
