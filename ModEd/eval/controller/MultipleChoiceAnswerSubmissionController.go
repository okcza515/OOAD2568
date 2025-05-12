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

func (c *MultipleChoiceAnswerSubmissionController) Grade(submissionID uint) (float64, error) {
	var score float64

	mcAnsSubs, err := c.List(map[string]interface{}{"submission_id": submissionID}, "Question")
	if err != nil {
		return 0, err
	}
	for _, mcAnsSub := range mcAnsSubs {
		if mcAnsSub.Choice.IsExpected {
			score += mcAnsSub.Question.Score
		}
	}
	return score, nil
}
