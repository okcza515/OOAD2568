package controller

import (
	"ModEd/core"
	"ModEd/project/model"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type PresentationController struct {
	*core.BaseController[*model.Presentation]
	db *gorm.DB
}

func NewPresentationController(db *gorm.DB) *PresentationController {
	return &PresentationController{
		db:             db,
		BaseController: core.NewBaseController[*model.Presentation](db),
	}
}

func (c *PresentationController) ListAllPresentations() ([]*model.Presentation, error) {
	return c.List(map[string]interface{}{})
}

func (c *PresentationController) RetrievePresentation(id uint) (*model.Presentation, error) {
	return c.RetrieveByID(id)
}

func (c *PresentationController) RetrievePresentationsBySeniorProjectId(projectId uint) ([]*model.Presentation, error) {
	return c.List(map[string]interface{}{"senior_project_id": projectId})
}

func (c *PresentationController) InsertPresentation(projectId uint, pType model.PresentationType, date time.Time) (*model.Presentation, error) {
	if !pType.IsValid() {
		return nil, fmt.Errorf("invalid presentation type")
	}
	if date.IsZero() {
		return nil, fmt.Errorf("presentation date cannot be empty")
	}

	presentation := &model.Presentation{
		PresentationType: pType,
		Date:             date,
		SeniorProjectId:  projectId,
	}

	return presentation, c.Insert(presentation)
}

func (c *PresentationController) UpdatePresentation(presentation *model.Presentation) error {
	if !presentation.PresentationType.IsValid() {
		return fmt.Errorf("invalid presentation type")
	}
	if presentation.Date.IsZero() {
		return fmt.Errorf("presentation date cannot be empty")
	}
	return c.UpdateByID(presentation)
}

func (c *PresentationController) DeletePresentation(id uint) error {
	return c.DeleteByID(id)
}
