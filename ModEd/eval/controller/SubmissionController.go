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

func (c *SubmissionController) GradingSubmission(submissionId uint) (submission *model.AnswerSubmission, err error) {
	var score = 0.0

	mcAnsSubs, err := c.McAnsSubController.GetMultipleChoiceAnswerSubmissionsBySubmissionID(submissionId)
	if err != nil {
		return nil, err
	}
	for _, mcAnsSub := range mcAnsSubs {
		if mcAnsSub.Choice.IsExpected {
			score += mcAnsSub.Question.Score
		}
	}

	tfAnsSubs, err := c.TfAnsSubController.GetTrueFalseAnswerSubmissionsBySubmissionID(submissionId)
	if err != nil {
		return nil, err
	}
	for _, tfAnsSub := range tfAnsSubs {
		tfAnswer, err := c.TfAnswerController.GetTrueFalseAnswerByQuestionID(tfAnsSub.QuestionID)
		if err != nil {
			return nil, err
		}

		if tfAnsSub.StudentAnswer == tfAnswer.IsExpected {
			score += tfAnsSub.Question.Score
		}
	}

	shortAnsSubs, err := c.ShortAnsSubController.GetShortAnswerSubmissionsBySubmissionID(submissionId)
	if err != nil {
		return nil, err
	}
	for _, shortAnsSub := range shortAnsSubs {
		shortAnswer, err := c.ShortAnswerController.GetShortAnswerByQuestionID(shortAnsSub.QuestionID)
		if err != nil {
			return nil, err
		}

		if shortAnsSub.StudentAnswer == shortAnswer.ExpectedAnswer {
			score += shortAnsSub.Question.Score
		}
	}

	submission, err = c.BaseController.RetrieveByCondition(map[string]interface{}{"id": submissionId})
	if err != nil {
		return nil, err
	}
	submission.Score = score
	if err := c.BaseController.UpdateByCondition(map[string]interface{}{"id": submissionId}, submission); err != nil {
		return nil, err
	}
	return submission, nil
}
