// MEP-1007
package controller

import (
	"ModEd/core"
	"ModEd/eval/model"

	"gorm.io/gorm"
)

type MultipleChoiceAnswerController struct {
	db *gorm.DB
	*core.BaseController[*model.MultipleChoiceAnswer]
}

func NewMultipleChoiceAnswerController(db *gorm.DB) *MultipleChoiceAnswerController {
	return &MultipleChoiceAnswerController{
		db:             db,
		BaseController: core.NewBaseController[*model.MultipleChoiceAnswer](db),
	}
}
