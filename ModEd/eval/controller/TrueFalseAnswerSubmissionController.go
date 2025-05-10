// MEP-1007
package controller

import (
	"ModEd/core"
	"ModEd/eval/model"

	"gorm.io/gorm"
)

type TrueFalseAnswerSubmissionController struct {
	db *gorm.DB
	*core.BaseController[*model.TrueFalseAnswerSubmission]
}

func NewTrueFalseAnswerSubmissionController(db *gorm.DB) *TrueFalseAnswerSubmissionController {
	return &TrueFalseAnswerSubmissionController{
		db:             db,
		BaseController: core.NewBaseController[*model.TrueFalseAnswerSubmission](db),
	}
}
