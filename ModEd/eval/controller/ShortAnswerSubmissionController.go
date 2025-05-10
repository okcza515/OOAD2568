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

func (c *ShortAnswerSubmissionController) GetShortAnswerSubmissionsBySubmissionID(submissionID uint) (shortAnsSubs []*model.ShortAnswerSubmission, err error) {
	err = c.db.
		Where("submission_id = ?", submissionID).
		Preload("Question").
		Find(&shortAnsSubs).Error

	if err != nil {
		return nil, err
	}

	return shortAnsSubs, err
}
