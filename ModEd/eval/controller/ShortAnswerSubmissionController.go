// MEP-1007
package controller

import (
	"ModEd/core"
	"ModEd/eval/model"
	"fmt"

	"gorm.io/gorm"
)

type ShortAnswerSubmissionController struct {
	db *gorm.DB
	*core.BaseController[*model.ShortAnswerSubmission]
	SAAnswerController *ShortAnswerController
}

func NewShortAnswerSubmissionController(db *gorm.DB) *ShortAnswerSubmissionController {
	return &ShortAnswerSubmissionController{
		db:                 db,
		BaseController:     core.NewBaseController[*model.ShortAnswerSubmission](db),
		SAAnswerController: NewShortAnswerController(db),
	}
}

func (c *ShortAnswerSubmissionController) Grade(submissionID uint) (float64, error) {
	var score float64

	shortAnsSubs, err := c.List(map[string]interface{}{"submission_id": submissionID}, "Question")
	if err != nil {
		return 0, err
	}
	for _, shortAnsSub := range shortAnsSubs {
		shortAnswer, err := c.SAAnswerController.RetrieveByCondition(map[string]interface{}{"question_id": shortAnsSub.QuestionID})
		if err != nil {
			return 0, err
		}

		if shortAnsSub.StudentAnswer == shortAnswer.ExpectedAnswer {
			fmt.Printf("Short score+=%f\n", shortAnsSub.Question.Score)
			score += shortAnsSub.Question.Score
		}
	}
	return score, nil
}
