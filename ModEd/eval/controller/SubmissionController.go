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
	MCAnsSubController *MultipleChoiceAnswerSubmissionController
	TFAnsSubController *TrueFalseAnswerSubmissionController
	SAAnsSubController *ShortAnswerSubmissionController
}

func NewSubmissionController(db *gorm.DB) *SubmissionController {
	return &SubmissionController{
		db:                 db,
		BaseController:     core.NewBaseController[*model.AnswerSubmission](db),
		MCAnsSubController: NewMultipleChoiceAnswerSubmissionController(db),
		TFAnsSubController: NewTrueFalseAnswerSubmissionController(db),
		SAAnsSubController: NewShortAnswerSubmissionController(db),
	}
}
func (c *SubmissionController) GradingSubmission(submissionID uint) (*model.AnswerSubmission, error) {
	var totalScore float64

	context := &GradingContext{}

	strategies := []GradingStrategy{
		c.MCAnsSubController,
		c.TFAnsSubController,
		c.SAAnsSubController,
	}

	for _, strategy := range strategies {
		context.SetStrategy(strategy)
		score, err := context.Grade(submissionID)
		if err != nil {
			return nil, err
		}
		totalScore += score
	}

	submission, err := c.BaseController.RetrieveByCondition(map[string]interface{}{"id": submissionID})
	if err != nil {
		return nil, err
	}
	submission.Score = totalScore
	if err := c.BaseController.UpdateByCondition(map[string]interface{}{"id": submissionID}, submission); err != nil {
		return nil, err
	}
	return submission, nil
}
