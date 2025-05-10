// MEP-1007
package controller

import (
	"ModEd/core"
	"ModEd/eval/model"

	"gorm.io/gorm"
)

type SubmissionController struct {
	db *gorm.DB
	*core.BaseController[*model.AnswerSubmission]
	McAnsSubController    *MultipleChoiceAnswerSubmissionController
	TfAnsSubController    *TrueFalseAnswerSubmissionController
	ShortAnsSubController *ShortAnswerSubmissionController
	TfAnswerController    *TrueFalseAnswerController
	ShortAnswerController *ShortAnswerController
}

func NewSubmissionController(db *gorm.DB) *SubmissionController {
	return &SubmissionController{
		db:                    db,
		BaseController:        core.NewBaseController[*model.AnswerSubmission](db),
		McAnsSubController:    NewMultipleChoiceAnswerSubmissionController(db),
		TfAnsSubController:    NewTrueFalseAnswerSubmissionController(db),
		ShortAnsSubController: NewShortAnswerSubmissionController(db),
		TfAnswerController:    NewTrueFalseAnswerController(db),
		ShortAnswerController: NewShortAnswerController(db),
	}
}
