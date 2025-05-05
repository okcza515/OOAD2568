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
