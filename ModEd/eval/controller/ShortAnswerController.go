// MEP-1007
package controller

import (
	"ModEd/core"
	"ModEd/eval/model"

	"gorm.io/gorm"
)

type ShortAnswerController struct {
	db *gorm.DB
	*core.BaseController[*model.ShortAnswer]
}

func NewShortAnswerController(db *gorm.DB) *ShortAnswerController {
	return &ShortAnswerController{
		db:             db,
		BaseController: core.NewBaseController[*model.ShortAnswer](db),
	}
}