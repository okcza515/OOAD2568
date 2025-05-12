// MEP-1007
package controller

import (
	"ModEd/core"
	"ModEd/eval/model"

	"gorm.io/gorm"
)

type TrueFalseAnswerController struct {
	db *gorm.DB
	*core.BaseController[*model.TrueFalseAnswer]
}

func NewTrueFalseAnswerController(db *gorm.DB) *TrueFalseAnswerController {
	return &TrueFalseAnswerController{
		db:             db,
		BaseController: core.NewBaseController[*model.TrueFalseAnswer](db),
	}
}