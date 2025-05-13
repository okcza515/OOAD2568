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
	GradingContext *GradingContext
}

func NewSubmissionController(db *gorm.DB) *SubmissionController {
	context := NewGradingContext()
	context.AddGradingStrategy(model.MultipleChoiceQuestion, NewMultipleChoiceAnswerSubmissionController(db))
	context.AddGradingStrategy(model.TrueFalseQuestion, NewTrueFalseAnswerSubmissionController(db))
	context.AddGradingStrategy(model.ShortAnswerQuestion, NewShortAnswerSubmissionController(db))
	return &SubmissionController{
		db:             db,
		BaseController: core.NewBaseController[*model.AnswerSubmission](db),
		GradingContext: context,
	}
}
func (c *SubmissionController) GradingSubmission(submissionID uint) (*model.AnswerSubmission, error) {

	submission, err := c.BaseController.RetrieveByCondition(map[string]interface{}{"id": submissionID})

	if err != nil {
		return nil, err
	}

	totalScore, err := c.GradingContext.GradeAll(submissionID)

	if err != nil {
		return nil, err
	}

	submission.Score = totalScore
	if err := c.BaseController.UpdateByCondition(map[string]interface{}{"id": submissionID}, submission); err != nil {
		return nil, err
	}
	return submission, nil
}
