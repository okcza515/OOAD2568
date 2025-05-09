// MEP-1007
package controller

import (
	"ModEd/core"
	"ModEd/eval/model"

	"gorm.io/gorm"
)

type TrueFalseAnswerSubmissionController struct {
	db *gorm.DB
	core *core.BaseController[*model.TrueFalseAnswerSubmission]
}

type ITrueFalseAnswerSubmission interface {
	CreateTrueFalseAnswerSubmission(tfAnsSub *model.TrueFalseAnswerSubmission) (tfAnsSubId uint, err error)
	GetAllTrueFalseAnswerSubmissions(preloads ...string) (tfAnsSubs []*model.TrueFalseAnswerSubmission, err error)
	GetTrueFalseAnswerSubmission(tfAnsSubId uint, preload ...string) (tfAnsSub *model.TrueFalseAnswerSubmission, err error)
	GetTrueFalseAnswerSubmissionsBySubmissionID(submissionID uint) ([]model.TrueFalseAnswerSubmission, error) 
	UpdateTrueFalseAnswerSubmission(updatedTfAnsSub *model.TrueFalseAnswerSubmission) (tfAnsSub *model.TrueFalseAnswerSubmission, err error)
	DeleteTrueFalseAnswerSubmission(tfAnsSubId uint) (tfAnsSub *model.TrueFalseAnswerSubmission, err error)
}

func NewTrueFalseAnswerSubmissionController(db *gorm.DB) *TrueFalseAnswerSubmissionController {
	return &TrueFalseAnswerSubmissionController{
		db: db,
		core: core.NewBaseController[*model.TrueFalseAnswerSubmission](db),
	}
}

func (c *TrueFalseAnswerSubmissionController) CreateTrueFalseAnswerSubmission(tfAnsSub *model.TrueFalseAnswerSubmission) (tfAnsSubId uint, err error) {
	if err := c.core.Insert(tfAnsSub); err != nil {
		return 0, err
	}
	return tfAnsSub.ID, nil
}

func (c *TrueFalseAnswerSubmissionController) GetAllTrueFalseAnswerSubmissions(preloads ...string) (tfAnsSubs []*model.TrueFalseAnswerSubmission, err error) {
	tfAnsSubs,err = c.core.List(nil,preloads...)
	if err != nil {
		return nil, err
	}
	return tfAnsSubs, nil
}

func (c *TrueFalseAnswerSubmissionController) GetTrueFalseAnswerSubmission(tfAnsSubId uint, preload ...string) (tfAnsSub *model.TrueFalseAnswerSubmission, err error) {
	tfAnsSub, err = c.core.RetrieveByCondition(map[string]interface{}{"id": tfAnsSubId}, preload...)
	if err != nil {
		return nil, err
	}
	return tfAnsSub, nil
}

func (c *TrueFalseAnswerSubmissionController) GetTrueFalseAnswerSubmissionsBySubmissionID(submissionID uint) ([]model.TrueFalseAnswerSubmission, error) {
    var tfAnsSub []model.TrueFalseAnswerSubmission
    err := c.db.
        Where("submission_id = ?", submissionID).
		Preload("Question").
        Find(&tfAnsSub).Error

	if err != nil {
		return nil, err
	}

    return tfAnsSub, err
}

func (c *TrueFalseAnswerSubmissionController) UpdateTrueFalseAnswerSubmission(updatedTfAnsSub *model.TrueFalseAnswerSubmission) (tfAnsSub *model.TrueFalseAnswerSubmission, err error){
	tfAnsSub, err = c.core.RetrieveByCondition(map[string]interface{}{"id": updatedTfAnsSub.ID})
	if err != nil {
		return nil, err
	}
	tfAnsSub.QuestionID = updatedTfAnsSub.QuestionID
	tfAnsSub.SubmissionID = updatedTfAnsSub.SubmissionID
	tfAnsSub.StudentAnswer = updatedTfAnsSub.StudentAnswer
	if err := c.core.UpdateByCondition(map[string]interface{}{"id": updatedTfAnsSub.ID}, tfAnsSub); err != nil{
		return nil, err
	}
	return tfAnsSub, nil
} 

func (c *TrueFalseAnswerSubmissionController) DeleteTrueFalseAnswerSubmission(tfAnsSubId uint) (tfAnsSub *model.TrueFalseAnswerSubmission, err error) {
	tfAnsSub, err = c.core.RetrieveByCondition(map[string]interface{}{"id": tfAnsSubId})
	if err != nil {
		return nil, err
	}

	if err := c.core.DeleteByCondition(map[string]interface{}{"id": tfAnsSubId}); err != nil {
		return nil, err
	}
	return tfAnsSub, nil
}