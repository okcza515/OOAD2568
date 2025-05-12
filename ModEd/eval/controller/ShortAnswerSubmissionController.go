// MEP-1007
package controller

import (
	"ModEd/core"
	"ModEd/eval/model"

	"gorm.io/gorm"
)

type ShortAnswerSubmissionController struct {
	db *gorm.DB
	*core.BaseController[*model.ShortAnswerSubmission]
}

func NewShortAnswerSubmissionController(db *gorm.DB) *ShortAnswerSubmissionController {
	return &ShortAnswerSubmissionController{
		db:             db,
		BaseController: core.NewBaseController[*model.ShortAnswerSubmission](db),
	}
}