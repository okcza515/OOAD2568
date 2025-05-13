// MEP-1007
package controller

import (
	"ModEd/core"
	"ModEd/eval/model"
	"fmt"

	"gorm.io/gorm"
)

type TrueFalseAnswerSubmissionController struct {
	db *gorm.DB
	*core.BaseController[*model.TrueFalseAnswerSubmission]
	TFAnswerController *TrueFalseAnswerController
}

func NewTrueFalseAnswerSubmissionController(db *gorm.DB) *TrueFalseAnswerSubmissionController {
	return &TrueFalseAnswerSubmissionController{
		db:                 db,
		BaseController:     core.NewBaseController[*model.TrueFalseAnswerSubmission](db),
		TFAnswerController: NewTrueFalseAnswerController(db),
	}
}

func (c *TrueFalseAnswerSubmissionController) Grade(submissionID uint) (float64, error) {
	var score float64

	tfAnsSubs, err := c.List(map[string]interface{}{"submission_id": submissionID}, "Question")
	if err != nil {
		return 0, err
	}

	for _, tfAnsSub := range tfAnsSubs {
		tfAnswer, err := c.TFAnswerController.RetrieveByCondition(map[string]interface{}{"question_id": tfAnsSub.QuestionID})
		if err != nil {
			return 0, err
		}

		if tfAnsSub.StudentAnswer == tfAnswer.IsExpected {
			fmt.Printf("TF score+=%f\n", tfAnsSub.Question.Score)
			score += tfAnsSub.Question.Score
		}
	}
	return score, nil
}
