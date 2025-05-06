package controller

import (
	"ModEd/core"
	"ModEd/project/model"

	"gorm.io/gorm"
)

type SeniorProjectController struct {
	*core.BaseController[*model.SeniorProject]
	DB *gorm.DB
}

func NewSeniorProjectController(db *gorm.DB) *SeniorProjectController {
	return &SeniorProjectController{
		BaseController: core.NewBaseController[*model.SeniorProject](db),
		DB:             db,
	}
}
