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

func (c *MultipleChoiceAnswerSubmissionController) GetMultipleChoiceAnswerSubmissionsBySubmissionID(submissionID uint) (mcAnsSubs []*model.MultipleChoiceAnswerSubmission, err error) {
	err = c.db.
		Where("submission_id = ?", submissionID).
		Preload("Question").
		Preload("Choice").
		Find(&mcAnsSubs).Error

	if err != nil {
		return nil, err
	}

	return mcAnsSubs, err
}
