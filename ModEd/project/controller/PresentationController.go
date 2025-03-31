package controller

import (
	"ModEd/core"
	"ModEd/project/model"

	"gorm.io/gorm"
)

type PresentationController struct {
	*core.BaseController
	db *gorm.DB
}

func NewPresentationController(db *gorm.DB) *PresentationController {
	return &PresentationController{
		db:             db,
		BaseController: core.NewBaseController("presentations", db),
	}
}

func (c *PresentationController) ListAllPresentations() ([]core.RecordInterface, error) {
	return c.List(map[string]interface{}{})
}

func (c *PresentationController) RetrievePresentation(id uint) (*core.RecordInterface, error) {
	return c.RetrieveByID(id)
}

func (c *PresentationController) InsertPresentation(presentation model.Presentation) error {
	return c.Insert(&presentation)
}

func (c *PresentationController) UpdatePresentation(id uint, presentation *model.Presentation) error {
	return c.UpdateByID(presentation)
}

func (c *PresentationController) DeletePresentation(id uint) error {
	return c.DeleteByID(id)
}
