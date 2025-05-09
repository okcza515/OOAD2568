// MEP-1007
package controller

import (
	"ModEd/core"
	"ModEd/eval/model"

	"gorm.io/gorm"
)

type MultipleChoiceAnswerSubmissionController struct {
	db *gorm.DB
	core *core.BaseController[*model.MultipleChoiceAnswerSubmission]
}

type IMultipleChoiceAnswerSubmission interface {
	CreateMultipleChoiceAnswerSubmission(mcAnsSub *model.MultipleChoiceAnswerSubmission) (mcAnsSubId uint, err error)
	GetAllMultipleChoiceAnswerSubmissions(preloads ...string) (mcAnsSubs []*model.MultipleChoiceAnswerSubmission, err error)
	GetMultipleChoiceAnswerSubmission(mcAnsSubId uint, preload ...string) (mcAnsSub *model.MultipleChoiceAnswerSubmission, err error)
	GetMultipleChoiceAnswerSubmissionsBySubmissionID(submissionID uint) ([]model.MultipleChoiceAnswerSubmission, error) 
	UpdateMultipleChoiceAnswerSubmission(updateMcAnsSub *model.MultipleChoiceAnswerSubmission) (mcAnsSub *model.MultipleChoiceAnswerSubmission, err error)
	DeleteMultipleChoiceAnswerSubmission(mcAnsSubId uint) (mcAnsSub *model.MultipleChoiceAnswerSubmission, err error)
}

func NewMultipleChoiceAnswerSubmissionController(db *gorm.DB) *MultipleChoiceAnswerSubmissionController {
	return &MultipleChoiceAnswerSubmissionController{
		db: db,
		core: core.NewBaseController[*model.MultipleChoiceAnswerSubmission](db),
	}
}

func (c *MultipleChoiceAnswerSubmissionController) CreateMultipleChoiceAnswerSubmission(mcAnsSub *model.MultipleChoiceAnswerSubmission) (mcAnsSubId uint, err error) {
	if err := c.core.Insert(mcAnsSub); err != nil {
		return 0, err
	}
	return mcAnsSub.ID, nil
}

func (c *MultipleChoiceAnswerSubmissionController) GetAllMultipleChoiceAnswerSubmissions(preloads ...string) (mcAnsSubs []*model.MultipleChoiceAnswerSubmission, err error) {
	mcAnsSubs,err = c.core.List(nil,preloads...)
	if err != nil {
		return nil, err
	}
	return mcAnsSubs, nil
}

func (c *MultipleChoiceAnswerSubmissionController) GetMultipleChoiceAnswerSubmission(mcAnsSubId uint, preload ...string) (mcAnsSub *model.MultipleChoiceAnswerSubmission, err error) {
	mcAnsSub, err = c.core.RetrieveByCondition(map[string]interface{}{"id": mcAnsSubId}, preload...)
	if err != nil {
		return nil, err
	}
	return mcAnsSub, nil
}

func (c *MultipleChoiceAnswerSubmissionController) GetMultipleChoiceAnswerSubmissionsBySubmissionID(submissionID uint) ([]model.MultipleChoiceAnswerSubmission, error) {
    var mcAnsSubs []model.MultipleChoiceAnswerSubmission
    err := c.db.
        Where("submission_id = ?", submissionID).
		Preload("Question").
        Preload("Choice").
        Find(&mcAnsSubs).Error

	if err != nil {
		return nil, err
	}

    return mcAnsSubs, err
}


func (c *MultipleChoiceAnswerSubmissionController) UpdateMultipleChoiceAnswerSubmission(updateMcAnsSub *model.MultipleChoiceAnswerSubmission) (mcAnsSub *model.MultipleChoiceAnswerSubmission, err error){
	mcAnsSub, err = c.core.RetrieveByCondition(map[string]interface{}{"id": updateMcAnsSub.ID})
	if err != nil {
		return nil, err
	}
	mcAnsSub.QuestionID = updateMcAnsSub.QuestionID
	mcAnsSub.SubmissionID = updateMcAnsSub.SubmissionID
	mcAnsSub.ChoiceID = updateMcAnsSub.ChoiceID
	if err := c.core.UpdateByCondition(map[string]interface{}{"id": updateMcAnsSub.ID}, mcAnsSub); err != nil{
		return nil, err
	}
	return mcAnsSub, nil
} 

func (c *MultipleChoiceAnswerSubmissionController) DeleteMultipleChoiceAnswerSubmission(mcAnsSubId uint) (mcAnsSub *model.MultipleChoiceAnswerSubmission, err error) {
	mcAnsSub, err = c.core.RetrieveByCondition(map[string]interface{}{"id": mcAnsSubId})
	if err != nil {
		return nil, err
	}

	if err := c.core.DeleteByCondition(map[string]interface{}{"id": mcAnsSubId}); err != nil {
		return nil, err
	}
	return mcAnsSub, nil
}