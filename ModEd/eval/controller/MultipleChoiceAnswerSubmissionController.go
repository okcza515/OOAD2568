// MEP-1007
package controller

import (
	"ModEd/core"
	"ModEd/eval/model"

	"gorm.io/gorm"
)

type MultipleChoiceAnswerSubmissionController struct {
	db *gorm.DB
	*core.BaseController[*model.MultipleChoiceAnswerSubmission]
}

func NewMultipleChoiceAnswerSubmissionController(db *gorm.DB) *MultipleChoiceAnswerSubmissionController {
	return &MultipleChoiceAnswerSubmissionController{
		db:             db,
		BaseController: core.NewBaseController[*model.MultipleChoiceAnswerSubmission](db),
	}
}