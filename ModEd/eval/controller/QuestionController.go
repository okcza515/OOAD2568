// MEP-1007
package controller

import (
	"ModEd/core"
	"ModEd/eval/model"

	"gorm.io/gorm"
)

type QuestionController struct {
	db *gorm.DB
	*core.BaseController[*model.Question]
}

func NewQuestionController(db *gorm.DB) *QuestionController {
	return &QuestionController{
		db:             db,
		BaseController: core.NewBaseController[*model.Question](db),
	}
}
