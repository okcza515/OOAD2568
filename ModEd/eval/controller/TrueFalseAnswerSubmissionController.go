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

func (c *TrueFalseAnswerSubmissionController) GetTrueFalseAnswerSubmissionsBySubmissionID(submissionID uint) (tfAnsSubs []*model.TrueFalseAnswerSubmission, err error) {
	err = c.db.
		Where("submission_id = ?", submissionID).
		Preload("Question").
		Find(&tfAnsSubs).Error

	if err != nil {
		return nil, err
	}

	return tfAnsSubs, err
}
