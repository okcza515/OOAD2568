// MEP-1007
package controller

import (
	"ModEd/core"
	"ModEd/eval/model"

	"gorm.io/gorm"
)

type ShortAnswerSubmissionController struct {
	db *gorm.DB
	core *core.BaseController[*model.ShortAnswerSubmission]
}

type IShortAnswerSubmission interface {
	CreateShortAnswerSubmission(shortAnsSub *model.ShortAnswerSubmission) (shortAnsSubId uint, err error)
	GetAllShortAnswerSubmissions(preloads ...string) (shortAnsSubs []*model.ShortAnswerSubmission, err error)
	GetShortAnswerSubmission(shortAnsSubId uint, preload ...string) (shortAnsSub *model.ShortAnswerSubmission, err error)
	GetShortAnswerSubmissionsBySubmissionID(submissionID uint) ([]model.ShortAnswerSubmission, error) 
	UpdateShortAnswerSubmission(updatedShortAnsSub *model.ShortAnswerSubmission) (shortAnsSub *model.ShortAnswerSubmission, err error)
	DeleteShortAnswerSubmission(shortAnsSubId uint) (shortAnsSub *model.ShortAnswerSubmission, err error)
}

func NewShortAnswerSubmissionController(db *gorm.DB) *ShortAnswerSubmissionController {
	return &ShortAnswerSubmissionController{
		db: db,
		core: core.NewBaseController[*model.ShortAnswerSubmission](db),
	}
}

func (c *ShortAnswerSubmissionController) CreateShortAnswerSubmission(shortAnsSub *model.ShortAnswerSubmission) (shortAnsSubId uint, err error) {
	if err := c.core.Insert(shortAnsSub); err != nil {
		return 0, err
	}
	return shortAnsSub.ID, nil
}

func (c *ShortAnswerSubmissionController) GetAllShortAnswerSubmissions(preloads ...string) (shortAnsSubs []*model.ShortAnswerSubmission, err error) {
	shortAnsSubs,err = c.core.List(nil,preloads...)
	if err != nil {
		return nil, err
	}
	return shortAnsSubs, nil
}

func (c *ShortAnswerSubmissionController) GetShortAnswerSubmission(shortAnsSubId uint, preload ...string) (shortAnsSub *model.ShortAnswerSubmission, err error) {
	shortAnsSub, err = c.core.RetrieveByCondition(map[string]interface{}{"id": shortAnsSubId}, preload...)
	if err != nil {
		return nil, err
	}
	return shortAnsSub, nil
}

func (c *ShortAnswerSubmissionController) GetShortAnswerSubmissionsBySubmissionID(submissionID uint) ([]model.ShortAnswerSubmission, error) {
    var shortAnsSub []model.ShortAnswerSubmission
    err := c.db.
        Where("submission_id = ?", submissionID).
		Preload("Question").
        Find(&shortAnsSub).Error

	if err != nil {
		return nil, err
	}

    return shortAnsSub, err
}


func (c *ShortAnswerSubmissionController) UpdateShortAnswerSubmission(updatedShortAnsSub *model.ShortAnswerSubmission) (shortAnsSub *model.ShortAnswerSubmission, err error){
	shortAnsSub, err = c.core.RetrieveByCondition(map[string]interface{}{"id": updatedShortAnsSub.ID})
	if err != nil {
		return nil, err
	}
	shortAnsSub.QuestionID = updatedShortAnsSub.QuestionID
	shortAnsSub.SubmissionID = updatedShortAnsSub.SubmissionID
	shortAnsSub.StudentAnswer = updatedShortAnsSub.StudentAnswer
	if err := c.core.UpdateByCondition(map[string]interface{}{"id": updatedShortAnsSub.ID}, shortAnsSub); err != nil{
		return nil, err
	}
	return shortAnsSub, nil
} 

func (c *ShortAnswerSubmissionController) DeleteShortAnswerSubmission(shortAnsSubId uint) (shortAnsSub *model.ShortAnswerSubmission, err error) {
	shortAnsSub, err = c.core.RetrieveByCondition(map[string]interface{}{"id": shortAnsSubId})
	if err != nil {
		return nil, err
	}

	if err := c.core.DeleteByCondition(map[string]interface{}{"id": shortAnsSubId}); err != nil {
		return nil, err
	}
	return shortAnsSub, nil
}