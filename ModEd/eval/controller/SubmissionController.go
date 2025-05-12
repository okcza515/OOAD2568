// MEP-1007
package controller

import (
	"ModEd/core"
	"ModEd/eval/model"
	"fmt"

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

func (c *SubmissionController) GradingSubmission(submissionID uint) (submission *model.AnswerSubmission, err error) {
	var score = 0.0

	mcAnsSubs, err := c.McAnsSubController.List(map[string]interface{}{"submission_id": submissionID}, "Question", "Choice")
	if err != nil {
		return nil, err
	}
	for _, mcAnsSub := range mcAnsSubs {
		if mcAnsSub.Choice.IsExpected {
			fmt.Printf("MC score+=%f\n", mcAnsSub.Question.Score)
			score += mcAnsSub.Question.Score
		}
	}

	tfAnsSubs, err := c.TfAnsSubController.List(map[string]interface{}{"submission_id": submissionID}, "Question")
	if err != nil {
		return nil, err
	}
	for _, tfAnsSub := range tfAnsSubs {
		tfAnswer, err := c.TfAnswerController.RetrieveByCondition(map[string]interface{}{"question_id": tfAnsSub.QuestionID})
		if err != nil {
			return nil, err
		}

		if tfAnsSub.StudentAnswer == tfAnswer.IsExpected {
			fmt.Printf("TF score+=%f\n", tfAnsSub.Question.Score)
			score += tfAnsSub.Question.Score
		}
	}

	shortAnsSubs, err := c.ShortAnsSubController.List(map[string]interface{}{"submission_id": submissionID}, "Question")
	if err != nil {
		return nil, err
	}
	for _, shortAnsSub := range shortAnsSubs {
		shortAnswer, err := c.ShortAnswerController.RetrieveByCondition(map[string]interface{}{"question_id": shortAnsSub.QuestionID})
		if err != nil {
			return nil, err
		}

		if shortAnsSub.StudentAnswer == shortAnswer.ExpectedAnswer {
			fmt.Printf("Short score+=%f\n", shortAnsSub.Question.Score)
			score += shortAnsSub.Question.Score
		}
	}

	submission, err = c.BaseController.RetrieveByCondition(map[string]interface{}{"id": submissionID})
	if err != nil {
		return nil, err
	}
	submission.Score = score
	if err := c.BaseController.UpdateByCondition(map[string]interface{}{"id": submissionID}, submission); err != nil {
		return nil, err
	}
	return submission, nil
}
