// MEP-1007
package controller

import (
	"ModEd/core"
	"ModEd/eval/model"

	"gorm.io/gorm"
)

type SubmissionController struct {
	db 						*gorm.DB
	core 					*core.BaseController[*model.AnswerSubmission]
	McAnsSubController 		*MultipleChoiceAnswerSubmissionController
	TfAnsSubController 		*TrueFalseAnswerSubmissionController
	ShortAnsSubController	*ShortAnswerSubmissionController
	TfAnswerController		*TrueFalseAnswerController
	ShortAnswerController	*ShortAnswerController
}

type ISubmission interface {
	CreateSubmission(submission *model.AnswerSubmission) (submissionId uint, err error)
	GetAllSubmissions(preloads ...string) (submissions []*model.AnswerSubmission, err error)
	GetSubmission(submissionId uint, preload ...string) (submission *model.AnswerSubmission, err error)
	UpdateSubmission(updatedSubmission *model.AnswerSubmission) (submission *model.AnswerSubmission, err error)
	GradingSubmission(submissionId uint) (submission *model.AnswerSubmission, err error)
	DeleteSubmission(submissionId uint) (submission *model.AnswerSubmission, err error)
}

func NewSubmissionController(db *gorm.DB) *SubmissionController {
	return &SubmissionController{
		db: db,
		core: core.NewBaseController[*model.AnswerSubmission](db),
	}
}

func (c *SubmissionController) CreateSubmission(submission *model.AnswerSubmission) (submissionId uint, err error) {
	if err := c.core.Insert(submission); err != nil {
		return 0, err
	}
	return submissionId, nil
}

func (c *SubmissionController) GetAllSubmissions(preloads ...string) (submissions []*model.AnswerSubmission, err error) {
	submissions,err = c.core.List(nil,preloads...)
	if err != nil {
		return nil, err
	}
	return submissions, nil
}

func (c *SubmissionController) GetSubmission(submissionId uint, preload ...string) (submission *model.AnswerSubmission, err error) {
	submission, err = c.core.RetrieveByCondition(map[string]interface{}{"id": submissionId}, preload...)
	if err != nil {
		return nil, err
	}
	return submission, nil
}

func (c *SubmissionController) UpdateSubmission(updatedSubmission *model.AnswerSubmission) (submission *model.AnswerSubmission, err error){
	submission, err = c.core.RetrieveByCondition(map[string]interface{}{"id": updatedSubmission.ID})
	if err != nil {
		return nil, err
	}
	submission.StudentID = updatedSubmission.StudentID
	submission.ExamID = updatedSubmission.ExamID
	submission.Score = updatedSubmission.Score
	if err := c.core.UpdateByCondition(map[string]interface{}{"id": updatedSubmission.ID}, submission); err != nil{
		return nil, err
	}
	return submission, nil
} 

func (c *SubmissionController) GradingSubmission(submissionId uint) (submission *model.AnswerSubmission, err error) {
	var score = 0.0

	mcAnsSubs, err := c.McAnsSubController.GetMultipleChoiceAnswerSubmissionsBySubmissionID(submissionId)
	if err != nil {
		return nil, err
	}
	for _, mcAnsSub := range mcAnsSubs {
		if mcAnsSub.Choice.IsExpected == true {
			score += mcAnsSub.Choice.Question.Score
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

	submission, err = c.core.RetrieveByCondition(map[string]interface{}{"id": submissionId})
	if err != nil {
		return nil, err
	}
	submission.Score = score
	if err := c.core.UpdateByCondition(map[string]interface{}{"id": submissionId}, submission); err != nil{
		return nil, err
	}
	return submission, nil
}

func (c *SubmissionController) DeleteSubmission(submissionId uint) (submission *model.AnswerSubmission, err error) {
	submission, err = c.core.RetrieveByCondition(map[string]interface{}{"id": submissionId})
	if err != nil {
		return nil, err
	}

	if err := c.core.DeleteByCondition(map[string]interface{}{"id": submissionId}); err != nil {
		return nil, err
	}
	return submission, nil
}